FROM golang:alpine as krypto

WORKDIR /app
ENV GO111MODULES=on

COPY . .

RUN go build -o /go/bin/krypto -v .

FROM alpine

COPY --from=krypto /go/bin /pkg/bin

ENTRYPOINT [ "/pkg/bin/krypto" ]