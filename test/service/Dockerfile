FROM golang:alpine

# Set environmet variables for test
ENV TEST_HOST=test \
    GATEWAY_HOST=web

# Set environmet variables for GO
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Move to working directory /build
WORKDIR /build

# Copy the code into the container
COPY . .

# Build the application
RUN go build -o bin/freighter cmd/freighter/main.go

# Move to /dist directory as the place for resulting binary folder
WORKDIR /dist

# Copy binary from build to main folder
RUN cp /build/bin/freighter .

# Export necessary port
EXPOSE 8000

# Command to run when starting the container
CMD ["/dist/freighter"]
