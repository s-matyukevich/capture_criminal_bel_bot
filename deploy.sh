#!/bin/bash 

set -eu

environment=$1

id=$(cat /dev/urandom | tr -dc 'a-z' | fold -w 8 | head -n 1)
tag=${environment}-${id}
echo "Tag: ${tag}"
docker build . -f Dockerfile -t smatyukevich/capture-criminal-tg-bot:${tag}
docker push smatyukevich/capture-criminal-tg-bot:${tag}

#gcloud auth login
gcloud container clusters get-credentials dapamazhy-by --zone europe-west3-a --project dapamazhy-by

helm upgrade -f ./k8s/values-${environment}.yaml \
  --set-file config=./config/${environment}-secret.yaml \
  --set image.tag=${tag}  capture-criminal-tg-bot-${environment} \
  --namespace tg-${environment} ./k8s 