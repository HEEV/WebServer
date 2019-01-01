# Adapted from https://blog.golang.org/docker and
# https://medium.com/@chemidy/create-the-smallest-and-secured-golang-docker-image-based-on-scratch-4752223b7324

# TODO: Look into using volumes for data storage

# Step 1: Preparation

# Start from the lightweight Alpine Linux-based FROM golang:alpine. This requires a little more initial configuration
# but will save space in the image
FROM golang:1.11.4-alpine AS builder

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git gcc

COPY . $GOPATH/src/github.com/HEEV/WebServer
WORKDIR $GOPATH/src/github.com/HEEV/WebServer

# Fetch dependencies
RUN go get -d -v

# Build our binary
# Simple build
#RUN go build -o /go/bin/server
# Linux-specific, CGO-disabled, no cross-compiling support for smaller image size
RUN GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o /go/bin/server

# Step 2: Build the image
FROM scratch

# Copy built executable
COPY --from=builder /go/bin/server /go/bin/server

# Run the server
ENTRYPOINT ["/go/bin/server"]

# Expose the service on internal port 8080
EXPOSE 8080