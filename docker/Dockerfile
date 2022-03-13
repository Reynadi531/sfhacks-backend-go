FROM golang:1.18rc1-alpine

RUN apk update && \
    apk add make git bash curl openssl alpine-sdk --no-cache

WORKDIR /app

COPY . .

RUN make build

CMD ["./bin/main"]
