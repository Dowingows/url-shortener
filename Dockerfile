# Stage 1: Build
FROM golang:1.25-alpine AS builder

# Dependências básicas
RUN apk add --no-cache git build-base

# Diretório da aplicação
WORKDIR /app

# Copiar go.mod e go.sum primeiro (cache de dependências)
COPY go.mod go.sum ./
RUN go mod download

# Copiar o restante do código
COPY . .

# Build do binário
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o url-shortener main.go

# Stage 2: Runtime
FROM alpine:latest

# Dependências mínimas para rodar
RUN apk add --no-cache ca-certificates

WORKDIR /root/

# Copiar binário do stage anterior
COPY --from=builder /app/url-shortener .

# Expor porta
EXPOSE 8080

# Rodar o binário
CMD ["./url-shortener"]
