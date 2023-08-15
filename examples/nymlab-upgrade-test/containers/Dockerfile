ARG BASE_IMAGE
ARG VERSION
ARG IMAGE_SOURCE
FROM ${BASE_IMAGE}:${VERSION} AS source

FROM alpine:3.16

LABEL org.opencontainers.image.source=${IMAGE_SOURCE}

COPY --from=source /bin /usr/bin
COPY --from=source /lib /usr/lib

# Set up dependencies
ENV PACKAGES curl make bash jq sed

# Install minimum necessary dependencies
RUN apk add --no-cache $PACKAGES

WORKDIR /opt
