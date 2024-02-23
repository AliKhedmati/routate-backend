FROM golang:alpine AS builder
LABEL maintainer="Ali Khedmati"
# ENV
ENV CGO_ENABLED 0
ENV GOPATH /go
ENV GOCACHE /go-build
ENV GOOS linux
# Args
ARG APP_PORT
# Copy
COPY . /routate
WORKDIR /routate
# Install Dependencies
RUN --mount=type=cache,target=/go/pkg/mod/cache \
    go mod download
# Build the binary
RUN --mount=type=cache,target=/go/pkg/mod/cache \
    --mount=type=cache,target=/go-build \
    go build -o routate-app src/cmd/main.go

# Start a new stage from scratch
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /routate/routate-app /usr/local/bin
EXPOSE ${APP_PORT}
CMD ["/usr/local/bin/routate-app"]

