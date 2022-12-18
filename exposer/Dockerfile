FROM golang:1.19-alpine AS build-env

# Set up dependencies
ENV PACKAGES curl make git libc-dev bash gcc linux-headers

# Set working directory for the build
WORKDIR /usr/local/share/app

# Add source files
COPY . .

# Install minimum necessary dependencies, build Cosmos SDK, remove packages
RUN apk add --no-cache $PACKAGES && go build -mod readonly -o exposer ./...

# Final image
FROM alpine:3.16

# Install ca-certificates
RUN apk add --update ca-certificates jq bash curl
WORKDIR /usr/local/share/app

RUN ls /usr/bin

# Copy over binaries from the build-env
COPY --from=build-env /usr/local/share/app/exposer /usr/bin/exposer

EXPOSE 8081
