FROM golang:alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build
COPY . .
RUN go build -o bin/freighter cmd/freighter/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates

ENV TEST_HOST=test \
    GATEWAY_HOST=web

WORKDIR /dist
COPY --from=builder /build/bin/freighter .
EXPOSE 8000
CMD ["/dist/freighter"]
