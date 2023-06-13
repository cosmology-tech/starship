ARG BASE_IMAGE
ARG VERSION=1.5.1
FROM ${BASE_IMAGE}:${VERSION}

USER root

ENV PACKAGES curl sed jq
RUN apt update && apt -y install $PACKAGES
