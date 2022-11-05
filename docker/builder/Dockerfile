FROM golang:1.18-alpine3.16

# Set up dependencies
ENV PACKAGES curl unzip make git libusb-dev libc-dev bash gcc linux-headers eudev-dev python3 musl-dev build-base

# Install ca-certificates
RUN set -eux; apk add --no-cache ca-certificates build-base;

# Install minimum necessary dependencies
RUN apk add --no-cache $PACKAGES

# Set working directory for the build
WORKDIR /usr/local/app
