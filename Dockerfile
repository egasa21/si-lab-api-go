FROM golang:1.23.4 AS builder

# Set the working directory
WORKDIR /app

# Copy the Go modules files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server ./cmd/api/main.go

# Use a minimal image to run the binary
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the compiled binary from the builder
COPY --from=builder /app/server .

# Expose the port Heroku provides
ENV PORT=8080
EXPOSE $PORT

# Run the application
CMD ["./server"]
