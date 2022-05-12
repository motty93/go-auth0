## Get auth0 access token
```
$ go run main.go
```

### Get start
```
$ docker build -t go-auth0:latest .

$ docker container run -it -d --name go-auth0-container go-auth0:latest sh

$ docker container cp go-auth0-container:/go/src/app/tmp ./

$ docker container rm -f go-auth0-container

$ ./tmp/main
```

もしくは
```
$ ./scripts/run.sh

$ ./scripts/build_run.sh
```

## env
```
$ go version
go version go1.17 linux/amd64
```

※docker imageはgo version1.18
