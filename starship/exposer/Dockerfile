FROM golang:1.21-alpine3.17 AS builder

LABEL org.opencontainers.image.source="https://github.com/hyperweb-io/starship"

WORKDIR /usr/local/app

COPY go.mod .
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -mod=readonly -o build/ ./...

FROM ghcr.io/hyperweb-io/starship/base:latest

COPY --from=builder /usr/local/app/build/exposer /bin

WORKDIR /opt

ENTRYPOINT ["exposer"]
