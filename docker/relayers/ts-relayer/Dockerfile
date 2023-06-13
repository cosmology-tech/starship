ARG BASE_IMAGE
ARG VERSION=0.9.0
FROM ${BASE_IMAGE}

# Set up dependencies
ENV PACKAGES curl make bash jq sed

# Install minimum necessary dependencies
RUN apk add --no-cache $PACKAGES

ENV VERSION=${VERSION}

# Install the relayer
RUN npm i -g @confio/relayer@${VERSION}

WORKDIR /opt
