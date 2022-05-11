## Get auth0 access token
```
$ go run main.go
```

### Start container
```
$ docker build -t go-auth0:latest .

$ docker run -it --rm go-auth0:latest go run main.go
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
