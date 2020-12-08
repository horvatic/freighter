FROM golang:alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build
COPY go.mod .
RUN go mod download
COPY . .
RUN go build -o bin/freighter cmd/freighter/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates

ENV INIT_API_KEY=123 \
    LISTEN_PORT=8080

WORKDIR /dist
COPY --from=builder /build/bin/freighter .
EXPOSE 8080
CMD ["/dist/freighter"]
