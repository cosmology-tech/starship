ARG BASE_IMAGE
ARG VERSION
FROM ${BASE_IMAGE}:${VERSION}

LABEL org.opencontainers.image.source="https://github.com/cosmology-tech/starship"

# Set up dependencies
ENV PACKAGES curl make bash jq sed

# Install minimum necessary dependencies
RUN apk update && apk add --no-cache $PACKAGES

WORKDIR /app

ENTRYPOINT ["/app/packages/faucet/bin/cosmos-faucet-dist"]
