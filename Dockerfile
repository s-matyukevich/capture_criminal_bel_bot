FROM golang:1.14.4 as build-backend
WORKDIR $GOPATH/src/github.com/s-matyukevich/capture-criminal-tg-bot

COPY ./src ./src
COPY ./main.go .
COPY ./go.mod .
COPY ./go.sum .

RUN go install .

FROM ubuntu
WORKDIR /app
RUN apt-get update
RUN apt-get install ca-certificates -y

COPY --from=build-backend /go/bin/capture-criminal-tg-bot /app/app

ENTRYPOINT /app/app