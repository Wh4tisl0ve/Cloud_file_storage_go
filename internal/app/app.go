package app

import (
	"fmt"
	"net/http"

	"github.com/Wh4tisl0ve/Cloud_file_storage_go/config"
	handlers "github.com/Wh4tisl0ve/Cloud_file_storage_go/internal/controller/http/handlers/user"
	"github.com/Wh4tisl0ve/Cloud_file_storage_go/internal/repository"
	usecase "github.com/Wh4tisl0ve/Cloud_file_storage_go/internal/usecase/user"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"

	// "github.com/Wh4tisl0ve/Cloud_file_storage_go/internal/repository"
	"github.com/Wh4tisl0ve/Cloud_file_storage_go/pkg/logger"
	"github.com/Wh4tisl0ve/Cloud_file_storage_go/pkg/postgres"
)

func Run(cfg *config.Config) {
	// logger setup
	envConfig := cfg.Env
	logger := logger.SetupLogger(envConfig)

	// DB setup
	dbConfig := cfg.Database

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Name)

	postgres, err := postgres.New(dsn)
	if err != nil {
		logger.Error(err.Error())
	} else {
		logger.Info("✅ Connected to PostgreSQL successfully!")
	}
	defer postgres.Close()

	// repositories
	userRepo := repository.New(postgres)

	// use-case
	createUserUC := usecase.NewCreateUserUseCase(userRepo)

	// routing and server
	// todo move to other folder
	r := chi.NewRouter()

	// middlewares
	r.Use(render.SetContentType(render.ContentTypeJSON))

	// routes
	r.Route("/api", func(r chi.Router) {
		// public routes
		r.Group(func(r chi.Router) {
			r.Post("/auth/sign-up", handlers.NewSignUpHandler(createUserUC))
		})
		// Private Routes
		// Require Authentication
		// r.Group(func(r chi.Router) {
		// 	r.Use(AuthMiddleware)
		// 	r.Post("/manage", CreateAsset)
		// })
	})

	// custom handlers
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, map[string]string{
			"error": "route not found",
		})
	})

	r.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		render.Status(r, http.StatusMethodNotAllowed)
		render.JSON(w, r, map[string]string{
			"error": "method not allowed",
		})
	})

	// todo .env config
	http.ListenAndServe("localhost:8000", r)
}
