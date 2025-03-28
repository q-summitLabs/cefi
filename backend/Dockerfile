# Build Stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Cache module downloads
COPY go.mod go.sum ./
RUN go mod download

# Copy source code and build binary
COPY . .
RUN go build -o server ./cmd/server

# Run Stage
FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/server .

EXPOSE 8080

ENTRYPOINT ["./server"]
