package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/egasa21/si-lab-api-go/configs"
	"github.com/egasa21/si-lab-api-go/internal/database"
	"github.com/egasa21/si-lab-api-go/internal/handler"
	"github.com/egasa21/si-lab-api-go/internal/middlewares"
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
	practicumRepository := repository.NewPracticumRepository(db)
	practicumModuleRepository := repository.NewPracticumModuleRepository(db)
	practicumModuleContentRepository := repository.NewPracticumModuleContentRepository(db)
	practicumClassRepository := repository.NewPracticumClassRepository(db)

	// Initialize services
	studentService := service.NewStudentService(studentRepository)
	authService := service.NewAuthService(authRepository)
	practicumService := service.NewPracticumService(practicumRepository)
	practicumModuleService := service.NewPracticumModuleService(practicumModuleRepository)
	practicumModuleContentService := service.NewPracticumModuleContentService(practicumModuleContentRepository)
	practicumClassService := service.NewPracticumClassService(practicumClassRepository)

	// Initialize handlers
	studentHandler := handler.NewStudentHandler(studentService)
	authHandler := handler.NewAuthHandler(authService)
	practicumHandler := handler.NewPracticumHandler(practicumService)
	practicumModuleHandler := handler.NewPracticumModuleHandler(practicumModuleService)
	practicumModuleContentHandler := handler.NewPracticumModuleContentHandler(practicumModuleContentService)
	practicumClassHandler := handler.NewPracticumClassHandler(practicumClassService)

	// Initialize main router
	mux := http.NewServeMux()
	v1Router := http.NewServeMux()

	// student
	v1Router.Handle("GET /students", middlewares.AuthMiddleware(authService)(http.HandlerFunc(studentHandler.GetAllStudents)))
	v1Router.Handle("GET /students/{id}", middlewares.AuthMiddleware(authService)(http.HandlerFunc(studentHandler.GetStudentById)))
	v1Router.HandleFunc("POST /students", studentHandler.CreateStudent)

	// practicum
	v1Router.HandleFunc("GET /practicums", practicumHandler.GetAllPracticums)
	v1Router.HandleFunc("POST /practicums", practicumHandler.CreatePracticum)
	v1Router.HandleFunc("GET /practicums/{id}", practicumHandler.GetPracticumByID)

	// practicum module
	v1Router.HandleFunc("GET /practicums/{practicum_id}/modules", practicumModuleHandler.GetModulesByPracticumID)
	v1Router.HandleFunc("POST /practicum-modules", practicumModuleHandler.CreateModule)
	v1Router.HandleFunc("GET /practicum-modules/{id}", practicumModuleHandler.GetModuleByID)

	// practicum module content
	v1Router.HandleFunc("POST /practicum-module-contents", practicumModuleContentHandler.CreateContent)
	v1Router.HandleFunc("GET /practicum-module-contents/{id}", practicumModuleContentHandler.GetContentByID)
	v1Router.HandleFunc("GET /practicum-modules/{module_id}/contents", practicumModuleContentHandler.GetContentsByModuleID)

	// practicum class
	v1Router.HandleFunc("POST /practicum-classes", practicumClassHandler.CreateClass)
	v1Router.HandleFunc("GET /practicum-classes/{id}", practicumClassHandler.GetClassByID)
	v1Router.HandleFunc("GET /practicums/{practicum_id}/classes", practicumClassHandler.GetClassesByPracticumID)
	v1Router.HandleFunc("PUT /practicum-classes/{id}", practicumClassHandler.UpdateClass)
	v1Router.HandleFunc("DELETE /practicum-classes/{id}", practicumClassHandler.DeleteClass)

	// auth
	v1Router.HandleFunc("POST /auth/register", authHandler.Register)
	v1Router.HandleFunc("POST /auth/login", authHandler.Login)

	// v1
	mux.Handle("/v1/", http.StripPrefix("/v1", v1Router))

	// Wrap the router with middleware
	handlerWithMiddleware := wrapMiddleware(mux, Logger(logger))

	// Setup HTTP server
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.AppPort),
		Handler: handlerWithMiddleware,
	}

	return &Server{
		server: server,
		logger: logger,
		mux:    mux,
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
