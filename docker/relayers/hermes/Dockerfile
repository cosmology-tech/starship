ARG BASE_IMAGE
ARG VERSION
FROM ${BASE_IMAGE}:${VERSION}

LABEL org.opencontainers.image.source="https://github.com/cosmology-tech/starship"

USER root

ENV PACKAGES curl sed jq
RUN apt update && apt -y install $PACKAGES
