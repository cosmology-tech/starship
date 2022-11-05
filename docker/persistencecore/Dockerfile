ARG CHAIN_NAME
ARG CHAIN_BIN
ARG CHAIN_VERSION
ARG CODE_REF

FROM golang:1.18-alpine3.16 AS go-builder

ENV CHAIN_NAME=${CHAIN_NAME:-"persistenceCore"}
ENV CHAIN_BIN=${CHAIN_BIN:-"persistenceCore"}
ENV CHAIN_VERSION=${CHAIN_VERSION:-"3.2.0"}
ENV CODE_REF=${CODE_REF:-"https://github.com/persistenceOne/persistenceCore/archive/refs/tags"}

# Set up dependencies
ENV PACKAGES curl make git libusb-dev libc-dev bash gcc linux-headers eudev-dev python3

# Install ca-certificates
RUN set -eux; apk add --no-cache ca-certificates build-base;

# Install minimum necessary dependencies
RUN apk add --no-cache $PACKAGES

# See https://github.com/CosmWasm/wasmvm/releases
ADD https://github.com/CosmWasm/wasmvm/releases/download/v1.1.1/libwasmvm_muslc.aarch64.a /lib/libwasmvm_muslc.aarch64.a
ADD https://github.com/CosmWasm/wasmvm/releases/download/v1.1.1/libwasmvm_muslc.x86_64.a /lib/libwasmvm_muslc.x86_64.a
RUN sha256sum /lib/libwasmvm_muslc.aarch64.a | grep 9ecb037336bd56076573dc18c26631a9d2099a7f2b40dc04b6cae31ffb4c8f9a
RUN sha256sum /lib/libwasmvm_muslc.x86_64.a | grep 6e4de7ba9bad4ae9679c7f9ecf7e283dd0160e71567c6a7be6ae47c81ebe7f32

# Copy the library you want to the final location that will be found by the linker flag `-lwasmvm_muslc`
RUN cp /lib/libwasmvm_muslc.$(uname -m).a /lib/libwasmvm_muslc.a

# Set working directory for the build
WORKDIR /usr/local/app

# Add source files
ADD $CODE_REF/v$CHAIN_VERSION.zip /tmp
RUN cd /tmp && unzip v$CHAIN_VERSION.zip && mv /tmp/${CHAIN_NAME}-${CHAIN_VERSION}/* /usr/local/app

# Force it to use static lib (from above) not standard libgo_cosmwasm.so file
RUN LEDGER_ENABLED=false BUILD_TAGS="muslc linkstatic" make build
RUN echo "Ensuring binary is statically linked ..." \
  && (file /usr/local/app/build/persistenceCore | grep "statically linked")

# Install cosmovisor
RUN go install github.com/cosmos/cosmos-sdk/cosmovisor/cmd/cosmovisor@v1.0.0

FROM alpine:3.16

ENV CHAIN_NAME=${CHAIN_NAME:-"persistenceCore"}
ENV CHAIN_BIN=${CHAIN_BIN:-"persistenceCore"}
ENV CHAIN_VERSION=${CHAIN_VERSION:-"3.2.0"}

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
