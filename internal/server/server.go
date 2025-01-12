package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/egasa21/si-lab-api-go/configs"
	"github.com/egasa21/si-lab-api-go/internal/database"
	"github.com/egasa21/si-lab-api-go/internal/handler"
	"github.com/egasa21/si-lab-api-go/internal/repository"
	"github.com/egasa21/si-lab-api-go/internal/service"
	_ "github.com/lib/pq"
)

type Server struct {
	server *http.Server
}

func NewServer(cfg *configs.Config) *Server {
	// Initialize DB connection
	db, err := database.ConnectDB(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Initialize repositories
	studentRepository := repository.NewStudentRepository(db)
	authRepository := repository.NewAuthRepository(db)

	// Initialize services
	studentService := service.NewStudentService(studentRepository)
	authService := service.NewAuthService(authRepository)

	// Initialize handlers
	studentHandler := handler.NewStudentHandler(studentService)
	authHandler := handler.NewAuthHandler(authService)
	mux := http.NewServeMux()

	// Setup v1 sub-mux
	v1 := http.NewServeMux()

	v1.HandleFunc("GET /students", studentHandler.GetAllStudents)
	v1.HandleFunc("GET /students/{id}", studentHandler.GetStudentById)
	v1.HandleFunc("POST /students", studentHandler.CreateStudent)

	// auth
	v1.HandleFunc("POST /auth/register", authHandler.Register)
	v1.HandleFunc("POST /auth/login", authHandler.Login)

	// Register the v1 routes
	mux.Handle("/v1/", http.StripPrefix("/v1", v1))

	// Setup the HTTP server
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.AppPort),
		Handler: mux,
	}

	return &Server{server: server}
}

// Start the HTTP server
func (s *Server) Start() error {
	fmt.Println("Starting server on port:", s.server.Addr)
	if err := s.server.ListenAndServe(); err != nil {
		log.Fatalf("Failed to start the server: %v", err)
		return err
	}
	return nil
}
