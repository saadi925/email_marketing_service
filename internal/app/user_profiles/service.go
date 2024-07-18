package userprofiles

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/saadi925/email_marketing_api/internal/database"
)

type UserProfileService interface {
	createUserProfile(ctx context.Context, userID uuid.UUID, req createUserProfileRequest) (database.CreateUserProfileParams, error)
	updateUserProfile(ctx context.Context, userID uuid.UUID, req updateUserProfileRequest) error
	getUserProfile(ctx context.Context, userID uuid.UUID) (database.UserProfile, error)
}

type userProfileService struct {
	db *database.Queries
}

func NewUserProfileService(db *database.Queries) UserProfileService {
	return &userProfileService{db: db}
}

func (s *userProfileService) createUserProfile(ctx context.Context, userID uuid.UUID, req createUserProfileRequest) (database.CreateUserProfileParams, error) {
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

	if _, err := s.db.CreateUserProfile(ctx, profile); err != nil {
		return profile, err
	}

	return profile, nil
}

func (s *userProfileService) updateUserProfile(ctx context.Context, userID uuid.UUID, req updateUserProfileRequest) error {
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

	if _, err := s.db.UpdateUserProfile(ctx, params); err != nil {
		return err
	}

	return nil
}

func (s *userProfileService) getUserProfile(ctx context.Context, userID uuid.UUID) (database.UserProfile, error) {
	profile, err := s.db.GetUserProfileByUserID(ctx, userID)
	if err != nil {
		return database.UserProfile{}, err
	}
	return profile, nil
}
