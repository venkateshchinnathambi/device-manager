# --- Stage 1: Builder ---
FROM golang:1.22-bullseye AS builder

WORKDIR /app

# Install build dependencies
RUN apt-get update && apt-get install -y \
  build-essential \
  pkg-config \
  librdkafka-dev

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy application code
COPY . .

# Build with CGO enabled
RUN CGO_ENABLED=1 go build -o thermostat-server main.go


# --- Stage 2: Runtime (still Debian-based) ---
FROM debian:bullseye-slim

WORKDIR /app

# Install runtime dependencies
RUN apt-get update && apt-get install -y librdkafka1 && apt-get clean

COPY --from=builder /app/thermostat-server .

EXPOSE 7777

ENTRYPOINT ["./thermostat-server"]
