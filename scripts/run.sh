#!/bin/sh

docker container run -it -d --name go-auth0-container go-auth0:latest sh
docker container cp go-auth0-container:/go/src/app/tmp ./
docker container rm -f go-auth0-container
./tmp/main
