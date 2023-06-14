# TikTok Immersion Programme 2023 

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

1. Clone this repository to your local machine: git clone https://github.com/NandiniVermaNTU/Submission.git
2. Change into the project directory: cd Submission
3. Install the project dependencies: go mod download

## Running the Servers

To run the HTTP server and gRPC server locally, follow these steps:

1. Start the gRPC server: 
Certainly! Here's the README file in Markdown format:

markdown
Copy code
# My Awesome Project

Welcome to My Awesome Project! This repository contains a simple implementation of an HTTP server and a gRPC server in Go. The HTTP server exposes an endpoint for sending messages, while the gRPC server provides a method for pulling messages for a specific recipient.

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
git clone https://github.com/your-username/my-awesome-project.git

csharp
Copy code

2. Change into the project directory:
cd my-awesome-project

markdown
Copy code

3. Install the project dependencies:
go mod download

yaml
Copy code

## Running the Servers

To run the HTTP server and gRPC server locally, follow these steps:

1. Start the gRPC server: go run rpc/main.go
2. In a separate terminal, start the HTTP server: go run http/main.go


## Testing the Implementation

Once the servers are running, you can test the implementation using different client tools.

### Testing the HTTP Server

To test the HTTP server, you can use a client tool like cURL or Postman.

Using cURL:
curl -X POST http://localhost:8080/send -d '{"recipient": "Alice", "message": "Hello, Alice!"}'


Using Postman:
- Create a new POST request with the URL `http://localhost:8080/send`.
- Set the request body to `{"recipient": "Alice", "message": "Hello, Alice!"}`.

### Testing the gRPC Server

To test the gRPC server, you need a gRPC client application. You can use the provided `client/main.go` as a starting point and modify it according to your requirements.


## Docker Support

If you prefer to run the servers with Docker, make sure you have Docker installed on your machine.

1. Build and start the servers using Docker Compose: docker-compose up --build
2. The HTTP server will be accessible at `http://localhost:8080`, and the gRPC server will be accessible at `localhost:50051`.





