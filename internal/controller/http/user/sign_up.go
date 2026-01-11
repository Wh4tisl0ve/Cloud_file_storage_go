package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
)

type SignUpRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type SignUpResponse struct {
	Username string `json:"username"`
}

type RegisterUserUseCase interface {
	Execute(username, password string) error
}

func NewSignUpHandler(uc RegisterUserUseCase) http.HandlerFunc {
	// todo upgrade error message
	// todo add response struct
	return func(w http.ResponseWriter, r *http.Request) {
		var req SignUpRequest

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

		err := uc.Execute(req.Username, req.Password)
		if err != nil {
			// todo что-то
			fmt.Println(err.Error())

			return
		}

		render.Status(r, http.StatusCreated)
		render.JSON(w, r, SignUpResponse{
			Username: req.Username,
		})
	}
}
