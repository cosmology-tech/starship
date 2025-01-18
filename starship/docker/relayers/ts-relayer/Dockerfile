ARG BASE_IMAGE
ARG VERSION
FROM ${BASE_IMAGE}

LABEL org.opencontainers.image.source="https://github.com/hyperweb-io/starship"

# Set up dependencies
ENV PACKAGES curl make bash jq sed

# Install minimum necessary dependencies
RUN apk add --no-cache $PACKAGES

ENV VERSION=${VERSION}

# Install the relayer
RUN npm i -g @confio/relayer@${VERSION}

WORKDIR /opt
