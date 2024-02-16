FROM golang:alpine AS builder
LABEL maintainer="Ali Khedmati"
# Args
ARG APP_NAME
ARG APP_PORT
# Copy
COPY . /routate/${APP_NAME}
WORKDIR /routate/${APP_NAME}
# Install Dependencies
RUN go mod download
# Build the binary
RUN go build -o ${APP_NAME} ./src/cmd/main.go

# Start a new stage from scratch
FROM golang:alpine
WORKDIR /routate/${APP_NAME}
COPY --from=builder /routate/${APP_NAME} .
CMD ./${APP_NAME}
# Expose
EXPOSE ${APP_PORT}

