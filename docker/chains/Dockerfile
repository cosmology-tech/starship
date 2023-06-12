ARG BASE_IMAGE
FROM ${BASE_IMAGE} AS source

FROM alpine:3.16

LABEL org.opencontainers.image.source="https://github.com/cosmology-tech/starship"

COPY --from=source /bin /usr/bin
COPY --from=source /lib /usr/lib

# Set up dependencies
ENV PACKAGES curl make bash jq sed

# Install minimum necessary dependencies
RUN apk add --no-cache $PACKAGES

WORKDIR /opt
