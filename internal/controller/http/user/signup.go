package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
)

type CreateUserRequest struct {
	UserName string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type CreateUserResponse struct {
	UserName string `json:"username"`
}

// todo command
type CreateUserUseCase interface {
	Execute(username, password string) error
}

func NewSignUpHandler(uc CreateUserUseCase) http.HandlerFunc {
	// todo upgrade error message
	// todo add response struct
	return func(w http.ResponseWriter, r *http.Request) {
		var req CreateUserRequest

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, map[string]string{
				"error": "invalid request body",
			})

			return
		}

		if err := validator.New().Struct(req); err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, map[string]string{
				"error": "validation failed",
			})

			return
		}

		err := uc.Execute(req.UserName, req.Password)
		if err != nil {
			// todo что-то
			fmt.Println(err.Error())
		}

		render.Status(r, http.StatusCreated)
		render.JSON(w, r, CreateUserResponse{req.UserName})
	}
}
