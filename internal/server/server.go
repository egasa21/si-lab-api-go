package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/egasa21/si-lab-api-go/configs"
	"github.com/egasa21/si-lab-api-go/internal/database"
	"github.com/egasa21/si-lab-api-go/internal/handler"
	"github.com/egasa21/si-lab-api-go/internal/repository"
	"github.com/egasa21/si-lab-api-go/internal/service"
	"github.com/rs/zerolog"
)

type middleware func(http.Handler) http.Handler

type Server struct {
	server *http.Server
	logger zerolog.Logger
	mux    *http.ServeMux
}

func NewServer(cfg *configs.Config, logger zerolog.Logger) *Server {
	// Initialize DB connection
	db, err := database.ConnectDB(cfg)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to connect to database")
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

	// Initialize main router
	mux := http.NewServeMux()
	v1Router := http.NewServeMux()

	v1Router.HandleFunc("GET /students", studentHandler.GetAllStudents)
	v1Router.HandleFunc("GET /students/{id}", studentHandler.GetStudentById)
	v1Router.HandleFunc("POST /students", studentHandler.CreateStudent)

	// auth
	v1Router.HandleFunc("POST /auth/register", authHandler.Register)
	v1Router.HandleFunc("POST /auth/login", authHandler.Login)

	// v1
	mux.Handle("/v1/", http.StripPrefix("/v1", v1Router))

	// Wrap the router with middleware
	handlerWithMiddleware := wrapMiddleware(v1Router, Logger(logger))

	// Setup HTTP server
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.AppPort),
		Handler: handlerWithMiddleware,
	}

	return &Server{
		server: server,
		logger: logger,
		mux:    v1Router,
	}
}

// Start the HTTP server
func (s *Server) Start() error {
	s.logger.Info().Msgf("Starting server on port: %s", s.server.Addr)
	if err := s.server.ListenAndServe(); err != nil {
		s.logger.Fatal().Err(err).Msg("Failed to start the server")
		return err
	}
	return nil
}

// Logger middleware logs all HTTP requests
func Logger(logger zerolog.Logger) middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			startTime := time.Now()

			// Serve the request
			next.ServeHTTP(w, r)

			// Log the request details
			duration := time.Since(startTime)
			logger.Info().
				Str("method", r.Method).
				Str("path", r.URL.Path).
				Str("remote_ip", r.RemoteAddr).
				Dur("duration", duration).
				Msg("Request completed")
		})
	}
}

// Middleware chaining function
func wrapMiddleware(handler http.Handler, middlewares ...middleware) http.Handler {
	// Apply middlewares in reverse order
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}
	return handler
}
