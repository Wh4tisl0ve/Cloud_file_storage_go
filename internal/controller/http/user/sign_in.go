package controller

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
)

type SignInRequest struct {
	Username string
	Password string
}

type SignInResponse struct {
	Username string
}

type AuthorizeUserUseCase interface {
	Execute(username, password string) error
}

func NewSignInHandler(uс AuthorizeUserUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req SignInRequest

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

		err := uс.Execute(req.Username, req.Password)
		if err != nil {
			render.Status(r, http.StatusUnauthorized)
			render.JSON(w, r, map[string]string{
				"error": "failed authorization",
			})

			return
		}

		render.Status(r, http.StatusOK)
		render.JSON(w, r, SignInResponse{
			Username: req.Username,
		})
	}
}
