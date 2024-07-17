package userprofiles

type createUserProfileRequest struct {
	Email         string `json:"email" validate:"required,email"`
	FirstName     string `json:"first_name" validate:"required"`
	LastName      string `json:"last_name" validate:"required"`
	PhoneNumber   string `json:"phone_number"`
	CompanyName   string `json:"company_name" validate:"required"`
	Website       string `json:"website"`
	StreetAddress string `json:"street_address" validate:"required"`
	ZipCode       string `json:"zip_code" validate:"required"`
	City          string `json:"city" validate:"required"`
	Country       string `json:"country" validate:"required"`
}

type updateUserProfileRequest struct {
	Email         string `json:"email" validate:"required,email"`
	FirstName     string `json:"first_name" validate:"required"`
	LastName      string `json:"last_name" validate:"required"`
	PhoneNumber   string `json:"phone_number"`
	CompanyName   string `json:"company_name" validate:"required"`
	Website       string `json:"website"`
	StreetAddress string `json:"street_address" validate:"required"`
	ZipCode       string `json:"zip_code" validate:"required"`
	City          string `json:"city" validate:"required"`
	Country       string `json:"country" validate:"required"`
}
