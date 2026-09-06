# Build stage
FROM golang:1.26-alpine AS builder

WORKDIR /app

# Cache Go dependencies
COPY go.mod ./
RUN go mod download

# Copy source
COPY . .

# Build a static binary
RUN CGO_ENABLED=0 GOOS=linux go build -o calculator .

# Runtime stage
FROM alpine:3.22

WORKDIR /app

# Copy the compiled application
COPY --from=builder /app/calculator .
COPY --from=builder /app/static ./static

EXPOSE 8080

CMD ["./calculator"]