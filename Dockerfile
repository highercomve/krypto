FROM golang:alpine as krypto

WORKDIR /app
ENV GO111MODULES=on

COPY . .

RUN go build -o /go/bin/krypto -v .

ENTRYPOINT [ "/go/bin/krypto" ]
