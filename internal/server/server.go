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

// todo: implement the student payment
// todo: update the class enrollment by payment success

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

	allowedOrigins := []string{"http://localhost:5173"}

	// Initialize repositories
	studentRepository := repository.NewStudentRepository(db)
	authRepository := repository.NewAuthRepository(db)
	practicumRepository := repository.NewPracticumRepository(db)
	practicumModuleRepository := repository.NewPracticumModuleRepository(db)
	practicumModuleContentRepository := repository.NewPracticumModuleContentRepository(db)
	practicumClassRepository := repository.NewPracticumClassRepository(db)
	studentRegistrationRepository := repository.NewStudentRegistrationRepository(db)
	studentClassEnrollmentRepository := repository.NewStudentClassEnrollmentRepository(db)
	userPracticumProgressRepository := repository.NewUserPracticumProgressRepository(db)
	userPracticumCheckpointRepository := repository.NewUserPracticumCheckpointRepository(db)

	// Initialize services
	studentService := service.NewStudentService(studentRepository)
	authService := service.NewAuthService(authRepository)
	practicumService := service.NewPracticumService(practicumRepository)
	practicumModuleService := service.NewPracticumModuleService(practicumModuleRepository)
	practicumModuleContentService := service.NewPracticumModuleContentService(practicumModuleContentRepository)
	practicumClassService := service.NewPracticumClassService(practicumClassRepository)
	studentRegistrationService := service.NewStudentRegistrationService(studentRegistrationRepository)
	studentClassEnrollmentService := service.NewStudentClassEnrollmentService(studentClassEnrollmentRepository)
	userPracticumProgressService := service.NewUserPracticumProgressService(userPracticumProgressRepository)
	userPracticumCheckpointService := service.NewUserPracticumCheckpointService(userPracticumCheckpointRepository)
	studentDataService := service.NewStudentDataService(userPracticumCheckpointService, practicumService, practicumModuleService, practicumModuleContentService)

	// Initialize handlers
	studentHandler := handler.NewStudentHandler(studentService, studentDataService)
	authHandler := handler.NewAuthHandler(authService)
	practicumHandler := handler.NewPracticumHandler(practicumService)
	practicumModuleHandler := handler.NewPracticumModuleHandler(practicumModuleService)
	practicumModuleContentHandler := handler.NewPracticumModuleContentHandler(practicumModuleContentService)
	practicumClassHandler := handler.NewPracticumClassHandler(practicumClassService)
	studentRegistrationHandler := handler.NewStudentRegistrationHandler(studentRegistrationService)
	studentClassEnrollmentHandler := handler.NewStudentClassEnrollmentHandler(studentClassEnrollmentService)
	userPracticumProgressHandler := handler.NewUserPracticumProgressHandler(userPracticumProgressService)
	userPracticumCheckpointHandler := handler.NewUserPracticumCheckpointHandler(userPracticumCheckpointService)

	// Initialize main router
	mux := http.NewServeMux()
	v1Router := http.NewServeMux()

	// student
	v1Router.Handle("GET /students", middlewares.AuthMiddleware(authService)(http.HandlerFunc(studentHandler.GetAllStudents)))
	v1Router.Handle("GET /students/{id}", middlewares.AuthMiddleware(authService)(http.HandlerFunc(studentHandler.GetStudentById)))
	v1Router.HandleFunc("POST /students", studentHandler.CreateStudent)
	v1Router.HandleFunc("GET /students/{id}/activities", studentHandler.GetStudentPracticumActivities)

	// student registration
	v1Router.HandleFunc("POST /student-registrations", studentRegistrationHandler.RegisterStudent)
	v1Router.HandleFunc("GET /students/{student_id}/registrations", studentRegistrationHandler.GetRegistrationsByStudentID)
	v1Router.HandleFunc("GET /practicums/{practicum_id}/registrations", studentRegistrationHandler.GetRegistrationsByPracticumID)
	v1Router.HandleFunc("DELETE /student-registrations/{id}", studentRegistrationHandler.DeleteRegistration)

	// student class enrollment
	v1Router.HandleFunc("POST /student-class-enrollments", studentClassEnrollmentHandler.EnrollStudent)
	v1Router.HandleFunc("GET /students/{student_id}/class-enrollments", studentClassEnrollmentHandler.GetEnrollmentsByStudentID)
	v1Router.HandleFunc("GET /practicum-classes/{class_id}/enrollments", studentClassEnrollmentHandler.GetEnrollmentsByClassID)
	v1Router.HandleFunc("DELETE /student-class-enrollments/{id}", studentClassEnrollmentHandler.UnenrollStudent)

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

	// user practicum progress
	v1Router.HandleFunc("POST /user-practicum-progress", userPracticumProgressHandler.CreateProgress)
	v1Router.HandleFunc("GET /user-practicum-progress/{user_id}/{practicum_id}", userPracticumProgressHandler.GetProgress)
	v1Router.HandleFunc("PUT /user-practicum-progress/{id}", userPracticumProgressHandler.UpdateProgress)
	v1Router.HandleFunc("PUT /user-practicum-progress/{user_id}/{practicum_id}/complete", userPracticumProgressHandler.MarkAsCompleted)
	v1Router.HandleFunc("DELETE /user-practicum-progress/{id}", userPracticumProgressHandler.DeleteProgress)

	// user practicum checkpoint
	v1Router.HandleFunc("POST /user-practicum-checkpoints", userPracticumCheckpointHandler.CreateCheckpoint)
	v1Router.HandleFunc("GET /user-practicum-checkpoints/{user_id}", userPracticumCheckpointHandler.GetCheckpointByUser)
	v1Router.HandleFunc("GET /user-practicum-checkpoints/{user_id}/{practicum_id}", userPracticumCheckpointHandler.GetCheckpointByUserAndPracticum)
	v1Router.HandleFunc("PUT /user-practicum-checkpoints/{id}", userPracticumCheckpointHandler.UpdateCheckpoint)
	v1Router.HandleFunc("DELETE /user-practicum-checkpoints/{id}", userPracticumCheckpointHandler.DeleteCheckpoint)

	// auth
	v1Router.HandleFunc("POST /auth/register", authHandler.Register)
	v1Router.HandleFunc("POST /auth/login", authHandler.Login)

	v1Router.HandleFunc("/health", healthCheckHandler)

	// v1
	mux.Handle("/v1/", http.StripPrefix("/v1", v1Router))

	// Wrap the router with middleware
	handlerWithMiddleware := wrapMiddleware(mux, CORSMiddleware(allowedOrigins), Logger(logger))

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

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	// You can add more checks here if needed (e.g., database, external services)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

// Middleware chaining function
func wrapMiddleware(handler http.Handler, middlewares ...middleware) http.Handler {
	// Apply middlewares in reverse order
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}
	return handler
}
