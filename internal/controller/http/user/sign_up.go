package controller

import (
	"encoding/json"
	"net/http"

	"github.com/Wh4tisl0ve/Cloud_file_storage_go/internal/controller/http/errors"
	"github.com/Wh4tisl0ve/Cloud_file_storage_go/internal/domain"
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
	return func(w http.ResponseWriter, r *http.Request) {
		var req SignUpRequest

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
			switch err {
			case domain.ErrInvalidPassword, domain.ErrInvalidUsername:
				render.Render(w, r, errors.ErrBadRequest(err))
			case domain.ErrUserAlreadyExists:
				render.Render(w, r, errors.ErrConflict(err))
			default:
				render.Render(w, r, errors.ErrBadRequest(err))
			}
			
			return
		}

		render.Status(r, http.StatusCreated)
		render.JSON(w, r, SignUpResponse{Username: req.Username})
	}
}
