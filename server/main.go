package main

import (
	"fmt"
	"log"
	"time"

	"github.com/a9jain/coral-backend/controllers"
	"github.com/a9jain/coral-backend/db"
	"github.com/a9jain/coral-backend/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/pebbe/zmq4"
)

func main() {
	go runSubscriber("tcp://127.0.0.1:5555", "status")
	var s controllers.Status = controllers.C_PlusPlus
	fmt.Println("Current Status:", s)
	db.Init() // Initialize the database connection
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // React dev server
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	routes.RegisterCompilerRoutes(r)
	r.Run(":8080") // listen and serve on localhost:8080

}

func runSubscriber(address string, topic string) {
	// Create ZeroMQ subscriber socket
	subscriber, err := zmq4.NewSocket(zmq4.SUB)
	if err != nil {
		log.Fatalf("Failed to create socket: %v", err)
	}
	defer subscriber.Close()

	// Connect to the given address
	if err := subscriber.Connect(address); err != nil {
		log.Fatalf("Failed to connect to %s: %v", address, err)
	}
	fmt.Printf("[Subscriber] Connected to %s\n", address)

	// Subscribe to the specified topic
	if err := subscriber.SetSubscribe(topic); err != nil {
		log.Fatalf("Failed to subscribe to topic %s: %v", topic, err)
	}
	fmt.Printf("[Subscriber] Subscribed to topic: %s\n", topic)

	for {
		msg, err := subscriber.RecvMessage(0)
		if err != nil {
			log.Printf("Error receiving message: %v", err)
			continue
		}

		fmt.Println("got one!") // Notification that a message was received

		// Print message contents
		if len(msg) >= 2 {
			fmt.Printf("Topic: %s | Message: %s\n", msg[0], msg[1])
		} else {
			fmt.Printf("Raw message: %v\n", msg)
		}
		// if you receive error, get the error message and find a way to send it to main backend function
	}

}
