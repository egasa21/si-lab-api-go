# Start with the Go base image
FROM golang:1.23.4 AS builder

# Set the working directory in the container
WORKDIR /app

# Copy the Go modules files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the Go application (disable CGO for compatibility)
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server ./cmd/api/main.go

# Use a minimal image to run the binary
FROM alpine:latest
RUN apk --no-cache add ca-certificates

# Set the working directory
WORKDIR /root/

# Copy the compiled binary from the builder
COPY --from=builder /app/server .

# Expose the port Heroku provides (port is defined as 8080 by default on Heroku)
ENV PORT=8080
EXPOSE 8080

# Command to run the application
CMD ["./server"]
