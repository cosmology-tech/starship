ARG CHAIN_NAME
ARG CHAIN_BIN
ARG CHAIN_VERSION
ARG CODE_REF

FROM golang:1.18-alpine3.16 AS go-builder

ENV CHAIN_NAME=${CHAIN_NAME:-"juno"}
ENV CHAIN_BIN=${CHAIN_BIN:-"junod"}
ENV CHAIN_VERSION=${CHAIN_VERSION:-"10.0.2"}
ENV CODE_REF=${CODE_REF:-"https://github.com/CosmosContracts/juno/archive/refs/tags"}

# Set up dependencies
ENV PACKAGES curl make git libusb-dev libc-dev bash gcc linux-headers eudev-dev python3

# Install ca-certificates
RUN set -eux; apk add --no-cache ca-certificates build-base;

# Install minimum necessary dependencies
RUN apk add --no-cache $PACKAGES

# See https://github.com/CosmWasm/wasmvm/releases
ADD https://github.com/CosmWasm/wasmvm/releases/download/v1.0.0/libwasmvm_muslc.aarch64.a /lib/libwasmvm_muslc.aarch64.a
ADD https://github.com/CosmWasm/wasmvm/releases/download/v1.0.0/libwasmvm_muslc.x86_64.a /lib/libwasmvm_muslc.x86_64.a
RUN sha256sum /lib/libwasmvm_muslc.aarch64.a | grep 7d2239e9f25e96d0d4daba982ce92367aacf0cbd95d2facb8442268f2b1cc1fc
RUN sha256sum /lib/libwasmvm_muslc.x86_64.a | grep f6282df732a13dec836cda1f399dd874b1e3163504dbd9607c6af915b2740479

# Copy the library you want to the final location that will be found by the linker flag `-lwasmvm_muslc`
RUN cp /lib/libwasmvm_muslc.$(uname -m).a /lib/libwasmvm_muslc.a

# Set working directory for the build
WORKDIR /usr/local/app

# Add source files
ADD $CODE_REF/v$CHAIN_VERSION.zip /tmp
RUN cd /tmp && unzip v${CHAIN_VERSION}.zip && mv /tmp/${CHAIN_NAME}-${CHAIN_VERSION}/* /usr/local/app

# Force it to use static lib (from above) not standard libgo_cosmwasm.so file
# force it to use static lib (from above) not standard libgo_cosmwasm.so file
# then log output of file /code/bin/junod
# then ensure static linking
RUN LEDGER_ENABLED=false BUILD_TAGS=muslc LINK_STATICALLY=true make build \
  && file /usr/local/app/bin/junod \
  && echo "Ensuring binary is statically linked ..." \
  && (file /usr/local/app/bin/junod | grep "statically linked")

FROM alpine:3.16

ENV CHAIN_NAME=${CHAIN_NAME:-"juno"}
ENV CHAIN_BIN=${CHAIN_BIN:-"junod"}
ENV CHAIN_VERSION=${CHAIN_VERSION:-"10.0.2"}

COPY --from=go-builder /usr/local/app/bin/$CHAIN_BIN /usr/bin/$CHAIN_BIN

# Set up dependencies
ENV PACKAGES curl make bash jq
ENV CHAIN_BIN /usr/bin/$CHAIN_BIN

# Install minimum necessary dependencies
RUN apk add --no-cache $PACKAGES

WORKDIR /opt

# rest server, p2p, rpc, grpc
EXPOSE 1317 26656 26657 9090

# Run persistenceCore by default, omit entrypoint to ease using container with cli
CMD ["/usr/bin/$CHAIN_BIN", "version", "--long"]
