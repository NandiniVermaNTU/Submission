# Use the official Golang image as the base image
FROM golang:1.16-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the necessary files to the working directory
COPY main.go ./
COPY messages.proto ./
COPY message/messages.pb.go ./message/
COPY message/messages_grpc.pb.go ./message/
COPY annotations.proto ./
COPY descriptor.proto ./
COPY http(1).proto ./

# Install protobuf compiler
RUN apk add --no-cache protobuf

# Generate protobuf code
RUN protoc --go_out=. messages.proto
RUN protoc --go_out=. --go-grpc_out=. messages.proto

# Build the Go application
RUN go build -o rpc-server

# Expose the port on which the server will listen
EXPOSE 50051

# Start the gRPC server
CMD ["./rpc-server"]
