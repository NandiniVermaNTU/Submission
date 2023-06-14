package main

import (
	"context"
	"log"
	"net"
	"sync"

	pb "mysubmission/rpc/message"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type messageServiceServer struct {
	pb.UnimplementedMessageServiceServer
	db *sql.DB // Added db variable
}

type PullMessagesResponse struct {
	Messages []*pb.Message `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
}

var (
	messages     []*pb.Message
	messagesLock sync.Mutex
	db           *sql.DB // Added db variable
)

func generateUniqueID() int32 {
	messagesLock.Lock()
	defer messagesLock.Unlock()

	return int32(len(messages) + 1)
}

func saveMessage(message *pb.Message, db *sql.DB) error {
	_, err := db.Exec(`
		INSERT INTO messages (sender, recipient, content)
		VALUES (?, ?, ?)
	`, message.Sender, message.Recipient, message.Content)
	if err != nil {
		log.Printf("Failed to save message: %v", err)
		return err
	}

	return nil
}

// Retrieve messages for the specified recipient from the messages table in the database
func filterMessagesByRecipient(receiver string, db *sql.DB) ([]*pb.Message, error) {
	rows, err := db.Query(`
		SELECT id, sender, recipient, content
		FROM messages
		WHERE recipient = ?
	`, receiver)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	filteredMessages := make([]*pb.Message, 0)
	for rows.Next() {
		var message pb.Message
		err := rows.Scan(&message.Id, &message.Sender, &message.Recipient, &message.Content)
		if err != nil {
			return nil, err
		}
		filteredMessages = append(filteredMessages, &message)
	}

	return filteredMessages, nil
}

// Store the message in the database
func storeMessage(db *sql.DB, sender, recipient, content string) error {
	stmt, err := db.Prepare("INSERT INTO messages (sender, recipient, content) VALUES (?, ?, ?)")
	if err != nil {
		log.Printf("Failed to prepare statement: %v", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(sender, recipient, content)
	if err != nil {
		log.Printf("Failed to store message: %v", err)
		return err
	}

	return nil
}

// Implement the Pull RPC method
func (s *messageServiceServer) Pull(ctx context.Context, req *pb.PullRequest) (*pb.PullResponse, error) {
	recipient := req.GetRecipient()

	filteredMessages, err := filterMessagesByRecipient(recipient, s.db)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to fetch messages: %v", err)
	}

	return &pb.PullResponse{Messages: filteredMessages}, nil
}

// Implement the SendMessage RPC method
func (s *messageServiceServer) SendMessage(ctx context.Context, req *pb.SendMessageRequest) (*pb.SendMessageResponse, error) {
	message := &pb.Message{
		Sender:    req.Sender,
		Recipient: req.Recipient,
		Content:   req.Content,
	}

	err := storeMessage(s.db, req.Sender, req.Recipient, req.Content)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to save message: %v", err)
	}

	response := &pb.SendMessageResponse{
		Id:        message.Id,
		Sender:    message.Sender,
		Recipient: message.Recipient,
		Content:   message.Content,
		Message:   message,
	}

	return response, nil
}

// Implement the FetchMessages RPC method
func (s *messageServiceServer) FetchMessages(ctx context.Context, req *pb.FetchMessagesRequest) (*pb.FetchMessagesResponse, error) {
	recipient := req.GetRecipient()

	filteredMessages, err := filterMessagesByRecipient(recipient, s.db)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to fetch messages: %v", err)
	}

	return &pb.FetchMessagesResponse{Messages: filteredMessages}, nil
}

func main() {
	// Open a connection to the MySQL database
	db, err := sql.Open("mysql", "root:tiktok@#123@tcp(localhost:3306)/Assignment_Database")
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	// Create the messages table if it doesn't exist
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS messages (
			id INT AUTO_INCREMENT PRIMARY KEY,
			sender VARCHAR(255) NOT NULL,
			recipient VARCHAR(255) NOT NULL,
			content TEXT,
			timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`)
	if err != nil {
		log.Fatalf("Failed to create messages table: %v", err)
	}

	// Set up the network binding
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Create a new gRPC server
	grpcServer := grpc.NewServer()

	// Register the RPC server implementation
	server := &messageServiceServer{
		db: db,
	}
	pb.RegisterMessageServiceServer(grpcServer, server)

	log.Println("RPC Server is running on http://localhost:50051")

	// Start the gRPC server
	log.Fatal(grpcServer.Serve(listen))
}
