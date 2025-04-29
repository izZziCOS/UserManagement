# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Download dependencies first (cacheable layer)
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build with optimizations
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o server ./cmd/server

# Runner stage (ultra-small image)
FROM gcr.io/distroless/static-debian12 AS runner

WORKDIR /app

# Copy binary and necessary files
COPY --from=builder /app/server .
COPY --from=builder /app/internal/config/config.yml ./internal/config/

# Non-root user for security
USER nonroot:nonroot

EXPOSE 8080
ENTRYPOINT ["./server"]