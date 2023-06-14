package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/mux"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Message struct {
	ID        int    `json:"id"`
	Sender    string `json:"sender"`
	Recipient string `json:"recipient"`
	Content   string `json:"content"`
}

type Messages struct {
	Data []Message `json:"data"`
}

type PullMessagesResponse struct {
	Messages []Message `json:"messages"`
}

var (
	messages         []Message
	messagesLock     sync.Mutex
	db               *sql.DB
	messageQueueSize = 100 // Adjust the queue size based on your requirements
)

// Generate a unique ID for the message
func generateUniqueID() int {
	messagesLock.Lock()
	defer messagesLock.Unlock()

	return len(messages) + 1
}

// Save the message to the messages table in the database
func storeMessage(db *sql.DB, sender, recipient, content string) error {
	_, err := db.Exec(`
		INSERT INTO messages (sender, recipient, content)
		VALUES (?, ?, ?)
	`, sender, recipient, content)
	if err != nil {
		return err
	}

	return nil
}

// Filter messages based on recipient
func filterMessagesByRecipient(receiver string, db *sql.DB) ([]Message, error) {
	messagesLock.Lock()
	defer messagesLock.Unlock()

	rows, err := db.Query(`
		SELECT id, sender, recipient, content
		FROM messages
		WHERE recipient = ?
	`, receiver)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	filteredMessages := make([]Message, 0)
	for rows.Next() {
		var message Message
		err := rows.Scan(&message.ID, &message.Sender, &message.Recipient, &message.Content)
		if err != nil {
			return nil, err
		}
		filteredMessages = append(filteredMessages, message)
	}

	return filteredMessages, nil
}

// Handle function for sending a message
func SendMessageHandler(w http.ResponseWriter, r *http.Request) {

	var message Message
	err := json.NewDecoder(r.Body).Decode(&message)

	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	message.ID = generateUniqueID()
	err = storeMessage(db, message.Sender, message.Recipient, message.Content)
	if err != nil {
		log.Printf("Failed to save message: %v", err)
		http.Error(w, "Failed to save message", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(message)
}

// Handle function for fetching messages for a recipient
func FetchMessagesHandler(w http.ResponseWriter, r *http.Request) {
	recipient := r.URL.Query().Get("recipient")
	messages, err := filterMessagesByRecipient(recipient, db)
	if err != nil {
		log.Printf("Failed to fetch messages: %v", err)
		http.Error(w, "Failed to fetch messages", http.StatusInternalServerError)
		return
	}

	// Add the retrieved messages to the message queue
	messagesLock.Lock()
	messages = append(messages, messages...)
	if len(messages) > messageQueueSize {
		messages = messages[len(messages)-messageQueueSize:]
	}
	messagesLock.Unlock()

	response := Messages{Data: messages}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// Handle function for pulling messages by receivers
func PullMessagesHandler(w http.ResponseWriter, r *http.Request) {
	recipient := r.URL.Query().Get("recipient")

	filteredMessages, err := filterMessagesByRecipient(recipient, db)
	if err != nil {
		log.Printf("Failed to fetch messages: %v", err)
		http.Error(w, "Failed to fetch messages", http.StatusInternalServerError)
		return
	}

	messagesLock.Lock()
	filteredMessages = append(filteredMessages, messages...)
	messagesLock.Unlock()

	response := PullMessagesResponse{
		Messages: filteredMessages,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
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

	router := mux.NewRouter()

	router.HandleFunc("/messages", SendMessageHandler).Methods(http.MethodPost)
	router.HandleFunc("/messages", FetchMessagesHandler).Methods(http.MethodGet).Queries("recipient", "{recipient}")
	router.HandleFunc("/pull", PullMessagesHandler).Methods(http.MethodGet)

	log.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
