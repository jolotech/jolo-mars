
# # =========================
# # 1) Build stage
# # =========================
# FROM golang:1.22-alpine AS builder

# WORKDIR /app

# # If you have private modules or need git for go mod
# RUN apk add --no-cache git ca-certificates

# # Copy go mod files first for better caching
# COPY go.mod go.sum ./
# RUN go mod download

# # Copy the rest of the code
# COPY . .

# # Build your app
# # CHANGE ./cmd/main.go if your entry file path is different
# RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server ./cmd/main.go mars


# # =========================
# # 2) Runtime stage
# # =========================
# FROM alpine:3.20

# WORKDIR /app

# # SSL certs for Stripe/HTTP calls
# RUN apk add --no-cache ca-certificates

# # Copy binary from builder
# COPY --from=builder /app/server /app/server

# # If you use .env inside container (optional)
# # COPY .env /app/.env

# # Your Go app port (change if different)
# EXPOSE 8023

# CMD ["/app/server"]



# =========================
# 1) Build stage
# =========================
FROM golang:1.25-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git ca-certificates

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# ✅ Build ONE main package only (remove that extra "mars")
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server ./cmd/main.go mars


# =========================
# 2) Runtime stage
# =========================
FROM alpine:3.20

WORKDIR /app
RUN apk add --no-cache ca-certificates

COPY --from=builder /app/server /app/server

EXPOSE 8023
CMD ["/app/server"]
