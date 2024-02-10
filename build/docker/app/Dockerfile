FROM golang:latest

LABEL maintainer="Ali Khedmati"

# Args
ARG APP_NAME
ARG APP_PORT

# Copy
COPY . /routate/${APP_NAME}
WORKDIR /routate/${APP_NAME}

# Install Dependencies
RUN go get ./

# Build the binary
RUN go build -o ${APP_NAME}

# Run the service
CMD ./${APP_NAME}

# Expose
EXPOSE ${APP_PORT}

