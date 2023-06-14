#TikTok Immersion Programme 2023 

Welcome to My Project! This repository contains a simple implementation of an HTTP server and a gRPC server in Go. The HTTP server exposes an endpoint for sending messages, while the gRPC server provides a method for pulling messages for a specific recipient.

## Directory Structure

- `http-server/`: Contains the source code for the HTTP server.
  - `main.go`: Entry point for the HTTP server.

- `rpc-server/`: Contains the source code for the gRPC server.
  - `main.go`: Entry point for the gRPC server.
  - `messages.proto`: Protocol Buffers definition file for message structures.
  - `message/`: Directory for generated message-related code.
    - `messages.pb.go`: Generated Go code for message structures.
    - `messages_grpc.pb.go`: Generated Go code for gRPC service.
  - `annotations.proto`: Protocol Buffers file for gRPC annotations.
  - `descriptor.proto`: Protocol Buffers file for service descriptors.
  - `http.proto`: Protocol Buffers file for HTTP annotations.

- `client/`: Contains the source code for a client application.
  - `main.go`: Entry point for the client application.

- `go.yml`: YAML file for the Go service configuration.

- `docker-compose.yml`: YAML file for Docker Compose configuration.

- `go.mod`, `go.sum`: Go module files.

## Prerequisites

Before running the servers, make sure you have the following prerequisites installed:

- Go (1.16 or higher)
- Docker (if you want to run the servers with Docker)
- cURL or Postman (for testing the HTTP server)
- gRPC client (for testing the gRPC server)

## Installation and Setup

1. Clone this repository to your local machine:
