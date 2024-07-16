package auth

type signInRequest struct {
	email    string `json:"email"`
	password string `json:"password"`
}

type signUpRequest struct {
	email    string `json:"email"`
	password string `json:"password"`
	name     string `json:"name"`
}
