ARG CHAIN_NAME
ARG CHAIN_BIN
ARG CHAIN_VERSION
ARG CODE_REF

FROM golang:1.18-alpine3.16 AS go-builder

ENV CHAIN_NAME=${CHAIN_NAME:-"gaia"}
ENV CHAIN_BIN=${CHAIN_BIN:-"gaiad"}
ENV CHAIN_VERSION=${CHAIN_VERSION:-"8.0.0"}
ENV CODE_REF=${CODE_REF:-"https://github.com/cosmos/gaia/archive/refs/tags"}

# Set up dependencies
ENV PACKAGES curl make git libusb-dev libc-dev bash gcc linux-headers eudev-dev python3

# Install ca-certificates
RUN set -eux; apk add --no-cache ca-certificates build-base;

# Install minimum necessary dependencies
RUN apk add --no-cache $PACKAGES

# Set working directory for the build
WORKDIR /usr/local/app

# Add source files
ADD $CODE_REF/v$CHAIN_VERSION.zip /tmp
RUN cd /tmp && unzip v$CHAIN_VERSION.zip && mv /tmp/${CHAIN_NAME}-${CHAIN_VERSION}/* /usr/local/app

# Force it to use static lib (from above) not standard libgo_cosmwasm.so file
RUN CGO_ENABLED=0 LEDGER_ENABLED=false make build

# Install cosmovisor
RUN go install github.com/cosmos/cosmos-sdk/cosmovisor/cmd/cosmovisor@v1.0.0

FROM alpine:3.16

ENV CHAIN_NAME=${CHAIN_NAME:-"gaia"}
ENV CHAIN_BIN=${CHAIN_BIN:-"gaiad"}
ENV CHAIN_VERSION=${CHAIN_VERSION:-"7.1.0"}

COPY --from=go-builder /usr/local/app/build/$CHAIN_BIN /usr/bin/$CHAIN_BIN
COPY --from=go-builder /go/bin/cosmovisor /usr/bin/cosmovisor

# Set up dependencies
ENV PACKAGES curl make bash jq sed
ENV CHAIN_BIN /usr/bin/$CHAIN_BIN

# Install minimum necessary dependencies
RUN apk add --no-cache $PACKAGES

WORKDIR /opt

# rest server, p2p, rpc, grpc
EXPOSE 1317 26656 26657 9090

# Run persistenceCore by default, omit entrypoint to ease using container with cli
CMD ["/usr/bin/$CHAIN_BIN", "version", "--long"]
