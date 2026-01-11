package app

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"

	"github.com/Wh4tisl0ve/Cloud_file_storage_go/internal/config"
	controller "github.com/Wh4tisl0ve/Cloud_file_storage_go/internal/controller/http/user"
	"github.com/Wh4tisl0ve/Cloud_file_storage_go/internal/service"
	storage "github.com/Wh4tisl0ve/Cloud_file_storage_go/internal/storage/postgres"
	"github.com/Wh4tisl0ve/Cloud_file_storage_go/internal/usecase"
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
	userRepo := storage.NewUserRepository(postgres)

	// services
	hasher := service.NewBcryptHasherService()

	// use-cases
	registerUserUC := usecase.NewRegisterUser(userRepo, hasher)
	authorizeUserUC := usecase.NewAuthorizeUser(userRepo, hasher)

	// routing and server
	// todo move to other folder
	r := chi.NewRouter()

	// middlewares
	r.Use(render.SetContentType(render.ContentTypeJSON))

	// routes
	r.Route("/api", func(r chi.Router) {
		// public routes
		r.Group(func(r chi.Router) {
			r.Post("/auth/sign-up", controller.NewSignUpHandler(registerUserUC))
		})
		r.Group(func(r chi.Router) {
			r.Post("/auth/sign-in", controller.NewSignInHandler(authorizeUserUC))
		})
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

	srvCfg := cfg.HttpServer

	addr := fmt.Sprintf("%s:%d", srvCfg.Host, srvCfg.Port)
	srv := &http.Server{Addr: addr, Handler: r}

	if err := srv.ListenAndServe(); err != nil {
		logger.Error(err.Error())
	}
}
