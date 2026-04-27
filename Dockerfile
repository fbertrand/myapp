# Dockerfile
FROM golang:1.24-alpine AS builder
WORKDIR /app

# Télécharger les dépendances en premier (cache Docker)
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build \
    -ldflags="-w -s -X main.version=$(git describe --tags --always)" \
    -o app ./cmd/main.go

# Image finale minimaliste
FROM scratch
COPY --from=builder /app/app /app
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

EXPOSE 8080
ENTRYPOINT ["/app"]