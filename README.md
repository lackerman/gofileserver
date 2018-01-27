# Basic static fileserver

## Build

- Build locally
```
go build
```

- Build docker image

```
docker build -t fileserver .
```

## Run server

- Run server locally
```
./gofileserver -dir . -port 8080
```

- Run docker image
```
docker run --rm -d -p 80:80 -v $PWD:/www fileserver ./go -port 80 -dir /www
```