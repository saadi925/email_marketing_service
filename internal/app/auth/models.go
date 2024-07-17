package auth

import (
	"github.com/go-playground/validator/v10"
)

type signUpRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
	Name     string `json:"name" validate:"required"`
}

type signInRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

var validate = validator.New()
