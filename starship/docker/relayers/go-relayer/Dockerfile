ARG BASE_IMAGE
ARG VERSION
FROM ${BASE_IMAGE}:${VERSION} as base

FROM alpine:3.17

LABEL org.opencontainers.image.source="https://github.com/hyperweb-io/starship"

# Set up dependencies
ENV PACKAGES --update ca-certificates curl make bash jq sed

# Install minimum necessary dependencies
RUN apk add --no-cache $PACKAGES

# Install rly binary
COPY --from=base /bin/rly /bin

WORKDIR /opt
