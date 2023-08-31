FROM golang:1.19-alpine3.16 AS builder

LABEL org.opencontainers.image.source="https://github.com/cosmology-tech/starship"

WORKDIR /usr/local/app

COPY go.mod .
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -mod=readonly -o build/ ./...

FROM ghcr.io/cosmology-tech/starship/base:latest

COPY --from=builder /usr/local/app/build/exposer /bin

WORKDIR /opt

ENTRYPOINT ["exposer"]
