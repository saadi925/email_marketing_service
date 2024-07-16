package auth

import (
	"net/http"

	"github.com/saadi925/email_marketing_api/internal/app/utils"
	"github.com/saadi925/email_marketing_api/internal/database"
)

func signIn(_ *database.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var creds signInRequest
		err := utils.ParseJSON(r, &creds)
		if err != nil {
			utils.RespondError(w, http.StatusBadRequest, "invalid request")
			return
		}
	}
}

func signUp(queries *database.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var creds signUpRequest
		err := utils.ParseJSON(r, &creds)
		if err != nil {
			utils.RespondError(w, http.StatusBadRequest, "invalid request")
			return
		}
		_, err = queries.CreateUser(r.Context(), database.CreateUserParams{
			Name:     creds.name,
			Password: creds.password,
			Email:    creds.email,
		})
		if err != nil {
			utils.RespondError(w, 500, "error occurred while saving user")
			return
		}

		utils.RespondJSON(w, 201, utils.MessageResponse{
			Message: "user has been created successfully",
		})

	}
}
func verifyEmail(_ *database.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
func resetPassword(_ *database.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
func verifyOneTimeToken(_ *database.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
