ARG BASE_IMAGE
ARG VERSION
FROM ${BASE_IMAGE}:${VERSION} as base

FROM alpine:3.16

LABEL org.opencontainers.image.source="https://github.com/cosmology-tech/starship"

# Set up dependencies
ENV PACKAGES --update ca-certificates curl make bash jq sed

# Install minimum necessary dependencies
RUN apk add --no-cache $PACKAGES

# Install rly binary
COPY --from=base /bin/rly /bin

WORKDIR /opt
