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
	db := database.ConnectDB(cfg)

	// Initialize repositories
	studentRepository := repository.NewStudentRepository(db)

	// Initialize services
	studentService := service.NewStudentService(studentRepository)

	// Initialize handlers
	studentHandler := handler.NewStudentHandler(studentService)

	mux := http.NewServeMux()

	// Setup v1 sub-mux
	v1 := http.NewServeMux()

	// Combined handler for /students to distinguish between GET and POST
	v1.HandleFunc("/students", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			studentHandler.GetAllStudents(w, r)
		case http.MethodPost:
			studentHandler.CreateStudent(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Specific handler for getting a student by ID
	v1.HandleFunc("/students/", studentHandler.GetStudentById)

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
