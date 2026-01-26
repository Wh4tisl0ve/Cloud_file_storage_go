package controller

import (
	"encoding/json"
	"net/http"

	"github.com/Wh4tisl0ve/Cloud_file_storage_go/internal/controller/http/errors"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
)

type SignInRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type SignInResponse struct {
	Username string `json:"username"`
}

type AuthorizeUserUseCase interface {
	Execute(username, password string) error
}

func NewSignInHandler(uc AuthorizeUserUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req SignInRequest

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			render.Render(w, r, errors.ErrBadRequest(err))

			return
		}

		if err := validator.New().Struct(req); err != nil {
			validateErrs := err.(validator.ValidationErrors)
			render.Render(w, r, errors.ErrValidation(validateErrs))

			return
		}
		
		if err := uc.Execute(req.Username, req.Password); err != nil {
			render.Render(w, r, errors.ErrUnauthorized(err))

			return
		}

		render.JSON(w, r, SignInResponse{Username: req.Username})
	}
}
