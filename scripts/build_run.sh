#!/bin/sh

docker build -t go-auth0:latest .
docker run -it --rm go-auth0:latest go run main.go
