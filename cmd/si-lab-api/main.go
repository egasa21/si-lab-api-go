package main

import (
	"fmt"
	"github.com/egasa21/si-lab-api-go/configs"
	"github.com/egasa21/si-lab-api-go/internal/server"
	"log"
)

func main() {
	// Load environment variables
	cfg := configs.LoadConfig()

	// Initialize the server
	srv := server.NewServer(cfg)

	// Start the server
	if err := srv.Start(); err != nil {
		log.Fatalf("Error starting server: %v", err)
		//os.Exit(1)
	}

	fmt.Println("Server started successfully")
}
