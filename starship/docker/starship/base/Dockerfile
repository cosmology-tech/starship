FROM alpine:3.17

LABEL org.opencontainers.image.source="https://github.com/hyperweb-io/starship"

# Set up dependencies
ENV PACKAGES --update ca-certificates curl make bash jq sed

# Install minimum necessary dependencies
RUN apk add --no-cache $PACKAGES

WORKDIR /opt
