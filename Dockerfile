# 2022-07-12 16:22:25.970205 +0200 CEST m=+0.159861209
# Create build stage based on buster image
FROM golang:1.19-rc-alpine3.15 AS builder
# Create working directory under /app
WORKDIR /app
# Copy over all go config (go.mod, go.sum etc.)
COPY go.* ./
# Install any required modules
RUN go mod download
# Copy over Go source code
COPY *.go ./
# Run the Go build and output binary under hello_go_http
RUN go build -o /hello_world
# Make sure to expose the port the HTTP server is using
# Run the app binary when we run the container


FROM alpine:3.16
COPY --from=builder /hello_world .
ENTRYPOINT ["/hello_world"]
