# Build Stage (Fast & Lightweight)
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Install git for any necessary dependencies
RUN apk add --no-cache git

# Copy go modules
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build a statically linked binary for maximum portability
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o vex .

# Final Stage (Minimal Image - perfect for BlackArch/Kali packaging later)
FROM alpine:latest

# Add ca-certificates in case we need to make HTTPS requests 
RUN apk --no-cache add ca-certificates

WORKDIR /app

# Bring the binary from the builder
COPY --from=builder /app/vex /usr/local/bin/vex

# Expose port if the mock server is running, though Vex is primarily a CLI
EXPOSE 8080

ENTRYPOINT ["vex"]
CMD ["--help"]
