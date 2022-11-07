package auth

type LoginResponse struct {
	Status     string     `json:"status"`
	Credential Credential `json:"data"`
}

type Credential struct {
	Token string `json:"token"`
}
