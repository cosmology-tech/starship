FROM confio/faucet:0.29.0

# Set up dependencies
ENV PACKAGES curl make bash jq sed

# Install minimum necessary dependencies
RUN apk update && apk add --no-cache $PACKAGES

WORKDIR /app

ENTRYPOINT ["/app/packages/faucet/bin/cosmos-faucet-dist"]
