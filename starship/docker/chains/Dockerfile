ARG BASE_IMAGE
ARG VERSION
FROM ${BASE_IMAGE}:${VERSION} AS source

FROM alpine:3.17

LABEL org.opencontainers.image.source="https://github.com/hyperweb-io/starship"

COPY --from=source /bin /usr/bin
COPY --from=source /lib /usr/lib

# Set up dependencies
ENV PACKAGES curl make bash jq sed

# Install minimum necessary dependencies
RUN apk add --no-cache $PACKAGES

WORKDIR /opt
