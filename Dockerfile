FROM golang:1.14.4 as build-backend
WORKDIR $GOPATH/src/github.com/s-matyukevich/capture-criminal-tg-bot

RUN apt-get update
RUN DEBIAN_FRONTEND=noninteractive apt-get install make git zlib1g-dev libssl-dev gperf php cmake g++ python -y
RUN git clone https://github.com/tdlib/td.git
RUN cd td && git checkout v1.6.0
RUN cd td && mkdir build && cd build && cmake -DCMAKE_BUILD_TYPE=Release ..
RUN cd td/build && cmake --build . -- -j8
RUN cd td/build && make install

COPY ./src ./src
COPY ./main.go .
COPY ./go.mod .
COPY ./go.sum .

RUN go install .

FROM ubuntu
WORKDIR /app
RUN DEBIAN_FRONTEND=noninteractive apt-get update && apt-get install ca-certificates python3 python3-pip -y
RUN pip3 install pymorphy2
RUN ln -s /usr/bin/python3 /usr/bin/python

COPY --from=build-backend /go/bin/capture-criminal-tg-bot /app/app

ENTRYPOINT /app/app