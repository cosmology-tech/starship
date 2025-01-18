ARG BASE_IMAGE
ARG VERSION
FROM ${BASE_IMAGE} as builder

LABEL org.opencontainers.image.source="https://github.com/hyperweb-io/starship"

USER root

RUN apt update && apt -y install ca-certificates jq bash curl sed unzip -y

ENV CODE_NAME="neutron-query-relayer"
ENV VERSION=${VERSION:-"refs/heads/main"}
ENV CODE_REF=${CODE_REF:-"https://github.com/neutron-org/neutron-query-relayer/archive"}

ADD $CODE_REF/$VERSION.zip /tmp
RUN mkdir /app && cd /tmp && unzip ${VERSION##*/}.zip && mv /tmp/${CODE_NAME}-${VERSION##*/}/* /app

WORKDIR /app

RUN go mod download
RUN go build -a -o build/neutron_query_relayer ./cmd/neutron_query_relayer/*.go

FROM debian:buster

RUN apt update && apt install ca-certificates jq bash curl sed -y && apt-get clean

ADD ["https://github.com/CosmWasm/wasmvm/releases/download/v1.5.2/libwasmvm.x86_64.so","https://github.com/CosmWasm/wasmvm/releases/download/v1.5.2/libwasmvm.aarch64.so","/lib/"]

COPY --from=builder /app/run.sh .
COPY --from=builder /app/build/neutron_query_relayer /bin/

EXPOSE 9999

ENTRYPOINT ["neutron_query_relayer", "start"]
