package userprofiles

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/saadi925/email_marketing_api/internal/app/middlewares"
	"github.com/saadi925/email_marketing_api/internal/app/utils"
)

func createUserProfile(service UserProfileService) middlewares.AuthHandler {
	return func(w http.ResponseWriter, r *http.Request, u uuid.UUID) {
		var req createUserProfileRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			utils.RespondError(w, http.StatusBadRequest, "Invalid input")
			return
		}

		if err := utils.Validate.Struct(req); err != nil {
			utils.RespondError(w, http.StatusBadRequest, "Validation failed: "+err.Error())
			return
		}

		profile, err := service.createUserProfile(r.Context(), u, req)
		if err != nil {
			utils.RespondError(w, http.StatusInternalServerError, "Failed to create profile")
			return
		}

		utils.RespondJSON(w, http.StatusCreated, profile)
	}
}

func updateUserProfile(service UserProfileService) middlewares.AuthHandler {
	return func(w http.ResponseWriter, r *http.Request, u uuid.UUID) {
		var req updateUserProfileRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			utils.RespondError(w, http.StatusBadRequest, "Invalid input")
			return
		}

		if err := utils.Validate.Struct(req); err != nil {
			utils.RespondError(w, http.StatusBadRequest, "Validation failed: "+err.Error())
			return
		}

		err := service.updateUserProfile(r.Context(), u, req)
		if err != nil {
			utils.RespondError(w, http.StatusInternalServerError, "Failed to update profile")
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

func getUserProfile(service UserProfileService) middlewares.AuthHandler {
	return func(w http.ResponseWriter, r *http.Request, u uuid.UUID) {
		profile, err := service.getUserProfile(r.Context(), u)
		if err != nil {
			if err == sql.ErrNoRows {
				utils.RespondError(w, http.StatusNotFound, "Profile not found")
			} else {
				utils.RespondError(w, http.StatusInternalServerError, "Failed to retrieve profile")
			}
			return
		}

		utils.RespondJSON(w, http.StatusOK, profile)
	}
}
