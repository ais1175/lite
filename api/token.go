package api

type CreateTokenRequest struct {
	Identifier string `json:"identifier"`
	Type       string `json:"type"`
}

type CreateTokenResponse struct {
	Token string `json:"token"`
}
