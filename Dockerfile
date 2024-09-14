
# =========================
# 1) Build stage
# =========================
FROM golang:1.25-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git ca-certificates

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server ./cmd


# =========================
# 2) Runtime stage
# =========================
FROM alpine:3.20

WORKDIR /app
RUN apk add --no-cache ca-certificates

COPY --from=builder /app/server /app/server
COPY --from=builder /app/templates /app/templates
COPY --from=builder /app/static /app/static

EXPOSE 8023
CMD ["/app/server"]
