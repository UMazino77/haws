# Use a lightweight Go base image
FROM golang:latest

# Install GCC, SQLite, and required dependencies
RUN apt-get update && apt-get install -y gcc sqlite3 libsqlite3-dev

# Set the working directory
WORKDIR /app

# Copy the application files
COPY . .

# Build the Go application
RUN go build -o server main.go

# Expose the required port (Railway automatically assigns a port)
EXPOSE 8080

# Start the server
CMD ["./server"]
