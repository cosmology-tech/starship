FROM public.ecr.aws/l9h3g6c6/injective-core:v1.11.5 AS source

FROM alpine:3.16

COPY --from=source /usr/local/bin/injectived /usr/bin

USER root

# Setup injective deps
RUN apk add --no-cache libgcc libstdc++ ncurses-libs

# Set up dependencies
ENV PACKAGES curl make bash jq sed

# Install minimum necessary dependencies
RUN apk add --no-cache $PACKAGES

WORKDIR /opt
