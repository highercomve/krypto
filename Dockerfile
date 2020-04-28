FROM golang:alpine as krypto

WORKDIR /app
ENV GO111MODULES=on

COPY . .

RUN GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o /go/bin/krypto -v .

FROM alpine

COPY --from=krypto /go/bin /pkg/bin

ENTRYPOINT [ "/pkg/bin/krypto" ]