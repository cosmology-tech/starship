ARG CODE_VERSION
ARG CODE_REF

FROM node:16

ENV CODE_NAME="explorer"
ENV CODE_VERSION=${CODE_VERSION:-"master"}
ENV CODE_REF=${CODE_REF:-"https://github.com/ping-pub/explorer/archive/refs/heads"}

WORKDIR /home/explorer

RUN apt-get update -y && apt-get install  ca-certificates jq bash curl -y

# Add source files
ADD $CODE_REF/$CODE_VERSION.zip /tmp
RUN cd /tmp && unzip ${CODE_VERSION}.zip && mv /tmp/${CODE_NAME}-${CODE_VERSION}/* /home/explorer

RUN echo "*.js\nsrc/\n" > .eslintignore

RUN rm -rf node_modules
RUN yarn install --frozen-lockfile

EXPOSE 8080
ENTRYPOINT ["yarn", "serve"]
