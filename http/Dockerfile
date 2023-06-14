# Use the official Golang image as the base image
FROM golang:1.16-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the necessary files to the working directory
COPY main.go ./

# Build the Go application
RUN go build -o http-server

# Expose the port on which the server will listen
EXPOSE 8080

# Start the HTTP server
CMD ["./http-server"]
