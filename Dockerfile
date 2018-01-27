FROM lackerman/gobuilder:latest as builder
LABEL AUTHOR=lackerman

WORKDIR /go/src/github.com/lackerman/fileserver

COPY *.go .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM scratch
LABEL AUTHOR=lackerman

WORKDIR /root

COPY --from=builder /go/src/github.com/lackerman/fileserver/app .

CMD ["./app", "/serve"]