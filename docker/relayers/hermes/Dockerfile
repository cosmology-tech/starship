FROM informalsystems/hermes:1.1.0

USER root

ENV PACKAGES curl sed jq
RUN apt update && apt -y install $PACKAGES
