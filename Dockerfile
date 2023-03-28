FROM golang:1.18-alpine as buildbase

RUN apk add git build-base

WORKDIR /go/src/github.com/dl-nft-books/book-svc
COPY vendor .
COPY . .

RUN GOOS=linux go build  -o /usr/local/bin/book-svc /go/src/github.com/dl-nft-books/book-svc


FROM alpine:3.9

COPY --from=buildbase /usr/local/bin/book-svc /usr/local/bin/book-svc
RUN apk add --no-cache ca-certificates

ENTRYPOINT ["book-svc"]
