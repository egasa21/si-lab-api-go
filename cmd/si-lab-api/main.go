package main

import (
	"fmt"
	"log"

	"github.com/egasa21/si-lab-api-go/configs"
	"github.com/egasa21/si-lab-api-go/internal/server"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/lib/pq"
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
