# 1. Build The Binary
FROM golang:latest AS builder

ENV APP_NAME routate

WORKDIR /routate/app
RUN apt update && apt upgrade --no-cache

RUN go mod download
COPY src ./
RUN go build -o routate-backend

# 2. Run The Binary.
FROM scratch
COPY --from=builder /rayantec/amitis/amitis /usr/local/bin/amitis
CMD ["/usr/local/bin/amitis"]
EXPOSE 3847
