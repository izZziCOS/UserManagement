# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Download dependencies first (cacheable layer)
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .
COPY .env .

# Build with optimizations
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o server ./cmd/server

# Runner stage (ultra-small image)
FROM alpine:3.18 AS runner

WORKDIR /app

# Install dependencies (ca-certificates for HTTPS, libc6-compat for Alpine)
RUN apk add --no-cache ca-certificates libc6-compat

# Copy binary and env files
COPY --from=builder /app/server .
COPY --from=builder /app/.env .

# Non-root user for security
RUN addgroup -S appgroup && adduser -S appuser -G appgroup
USER appuser

EXPOSE 8080
ENTRYPOINT ["./server"]