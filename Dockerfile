# ===== Build stage =====
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Install git (sometimes needed for go mod)
RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build (adjust main path to your project)
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server ./cmd/main.go

# ===== Run stage =====
FROM alpine:3.20

WORKDIR /app

# certs for https calls (stripe, etc)
RUN apk add --no-cache ca-certificates

COPY --from=builder /app/server /app/server

EXPOSE 8023

# If your app needs migrations etc, do it in entrypoint script (optional)
CMD ["make start"]
