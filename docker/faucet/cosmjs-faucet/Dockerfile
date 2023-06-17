ARG BASE_IMAGE
ARG VERSION

FROM ${BASE_IMAGE} as build-env

ENV CODE_VERSION=${VERSION}
ENV CODE_REF="https://github.com/cosmos/cosmjs/archive/refs/tags"

LABEL org.opencontainers.image.source="https://github.com/cosmology-tech/starship"

# Set up dependencies
ENV PACKAGES curl make bash jq sed

# Install minimum necessary dependencies
RUN apk update && apk add --no-cache $PACKAGES

WORKDIR /app

# Fetch cosmjs from github
ADD $CODE_REF/$CODE_VERSION.tar.gz cosmjs.tar.gz
RUN tar --strip-components=1 -xvf cosmjs.tar.gz
RUN rm -rf cosmjs.tar.gz

# Install cosmjs faucet
RUN apk add --update --no-cache alpine-sdk linux-headers build-base gcc libusb-dev python3 py3-pip eudev-dev
RUN ln -sf python3 /usr/bin/python

RUN yarn install && yarn run build
RUN (cd packages/faucet && SKIP_BUILD=1 yarn pack-node)

# Use Alpine and install Node.js which is 50% smaller than the -alpine version of the node
# image (53 MB including the faucet app).
FROM alpine:3.15

# Set up dependencies, nodejs is special case
ENV PACKAGES curl make bash jq sed nodejs

# Install minimum necessary dependencies
RUN apk update && apk add --no-cache $PACKAGES

COPY --from=build-env /app/packages/faucet /app/packages/faucet
WORKDIR /app

EXPOSE 8000

ENTRYPOINT ["/app/packages/faucet/bin/cosmos-faucet-dist"]

