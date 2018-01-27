FROM golang:1.9.3 as builder

WORKDIR /go/src/github.com/lackerman/gofileserver

COPY *.go .

RUN go get -d -v golang.org/x/net/html && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .

FROM scratch
LABEL AUTHOR=lackerman

WORKDIR /bin

COPY --from=builder /go/src/github.com/lackerman/gofileserver/gofileserver .

CMD ["gofileserver"]