FROM golang:1.21-alpine3.17

LABEL org.opencontainers.image.source="https://github.com/hyperweb-io/starship"

# Set up dependencies
ENV PACKAGES curl unzip make git libusb-dev libc-dev bash gcc linux-headers eudev-dev python3 musl-dev

# Install ca-certificates
RUN set -eux; apk add --no-cache ca-certificates build-base;

# Install minimum necessary dependencies
RUN apk add --no-cache $PACKAGES

# Set working directory for the build
WORKDIR /usr/local/app
