ARG BASE_IMAGE
ARG VERSION

FROM ${BASE_IMAGE}

LABEL org.opencontainers.image.source="https://github.com/hyperweb-io/starship"

ENV CODE_NAME="explorer"
ENV VERSION=${VERSION:-"refs/heads/master"}
ENV CODE_REF=${CODE_REF:-"https://github.com/ping-pub/explorer/archive"}

WORKDIR /home/explorer

RUN apt-get update -y && apt-get install  ca-certificates jq bash curl -y

# Add source files
ADD $CODE_REF/$VERSION.zip /tmp
RUN cd /tmp && unzip ${VERSION##*/}.zip && mv /tmp/${CODE_NAME}-${VERSION##*/}/* /home/explorer

RUN echo "*.js\nsrc/\n" > .eslintignore

RUN rm -rf node_modules
RUN yarn install --no-lockfile --network-timeout 100000

EXPOSE 8080
ENTRYPOINT ["yarn", "serve"]
