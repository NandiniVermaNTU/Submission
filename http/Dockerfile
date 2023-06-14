# Using the official Golang image as the base image
FROM golang:1.16-alpine

# Setting the working directory inside the container
WORKDIR /app

# Copying the necessary files to the working directory
COPY main.go ./

# Building the Go application
RUN go build -o http-server

# Exposing the port on which the server will listen
EXPOSE 8080

# Starting the HTTP server
CMD ["./http-server"]
