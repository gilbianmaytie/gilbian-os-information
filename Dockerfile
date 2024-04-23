# Use the official Golang image as the base image
FROM golang:1.16 AS builder

# Set the current working directory inside the container
WORKDIR /app

# Copy the Go module files to the container
COPY go.mod .

# Download and install Go dependencies
RUN go mod download

# Copy the rest of the application source code to the container
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o gilbian-os-information .

# Start a new stage from scratch
FROM alpine:latest

# Set the current working directory inside the container
WORKDIR /root/

# Copy the built executable from the previous stage
COPY --from=builder /app/gilbian-os-information .

# Expose port 8080 (optional)
EXPOSE 8080

# Command to run the executable
CMD ["./gilbian-os-information"]
