# Build Stage
FROM golang:1.21

# Set the working directory
WORKDIR /app

# Copy the Go modules manifests
COPY go.mod go.sum ./

# Download Go modules
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o main .

# Expose the port that the app runs on
EXPOSE 8080

# Run the application
CMD ["./main"]
