package auth

type LoginResponse struct {
	Status     string     `json:"status"`
	Credential Credential `json:"data"`
}

type CredentialRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Credential struct {
	Token string `json:"token"`
}
