ARG GOLANG_VERSION="1.19"

FROM golang:${GOLANG_VERSION}-alpine AS builder

RUN apk --no-cache add \
        tzdata

WORKDIR /code

COPY . .

RUN go build -ldflags '-s' -o ./socks5

FROM alpine

COPY --from=builder /code/socks5 /

ENTRYPOINT ["/socks5"]
