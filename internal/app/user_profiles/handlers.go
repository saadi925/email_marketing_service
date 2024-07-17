package userprofiles

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/saadi925/email_marketing_api/internal/app/utils"
	"github.com/saadi925/email_marketing_api/internal/database"
)

func createUserProfile(w http.ResponseWriter, r *http.Request, userID uuid.UUID, db *database.Queries) {
	var req createUserProfileRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := utils.Validate.Struct(req); err != nil {
		http.Error(w, "Validation failed: "+err.Error(), http.StatusBadRequest)
		return
	}

	profile := database.CreateUserProfileParams{
		UserID:        userID,
		Email:         req.Email,
		FirstName:     req.FirstName,
		LastName:      req.LastName,
		PhoneNumber:   sql.NullString{String: req.PhoneNumber, Valid: req.PhoneNumber != ""},
		CompanyName:   req.CompanyName,
		Website:       sql.NullString{String: req.Website, Valid: req.Website != ""},
		StreetAddress: req.StreetAddress,
		ZipCode:       req.ZipCode,
		City:          req.City,
		Country:       req.Country,
	}

	if _, err := db.CreateUserProfile(r.Context(), profile); err != nil {
		http.Error(w, "Failed to create profile", http.StatusInternalServerError)
		return
	}
	utils.RespondJSON(w, http.StatusCreated, profile)
}

func updateUserProfile(w http.ResponseWriter, r *http.Request, userID uuid.UUID, db *database.Queries) {
	var req updateUserProfileRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := utils.Validate.Struct(req); err != nil {
		http.Error(w, "Validation failed: "+err.Error(), http.StatusBadRequest)
		return
	}

	params := database.UpdateUserProfileParams{
		UserID:        userID,
		Email:         req.Email,
		FirstName:     req.FirstName,
		LastName:      req.LastName,
		PhoneNumber:   sql.NullString{String: req.PhoneNumber, Valid: req.PhoneNumber != ""},
		CompanyName:   req.CompanyName,
		Website:       sql.NullString{String: req.Website, Valid: req.Website != ""},
		StreetAddress: req.StreetAddress,
		ZipCode:       req.ZipCode,
		City:          req.City,
		Country:       req.Country,
	}

	if _, err := db.UpdateUserProfile(r.Context(), params); err != nil {
		http.Error(w, "Failed to update profile", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func getUserProfile(w http.ResponseWriter, r *http.Request, userID uuid.UUID, db *database.Queries) {
	profile, err := db.GetUserProfileByUserID(r.Context(), userID)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Profile not found", http.StatusNotFound)
		} else {
			http.Error(w, "Failed to retrieve profile", http.StatusInternalServerError)
		}
		return
	}

	utils.RespondJSON(w, http.StatusOK, profile)
}
