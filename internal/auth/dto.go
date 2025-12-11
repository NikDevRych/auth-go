package auth

import "time"

type RefreshTokenResponse struct {
	RefreshToken string    `json:"token"`
	ExpireAt     time.Time `json:"expire_at"`
}

type TokenResponse struct {
	AccessToken  string               `json:"access_token"`
	RefreshToken RefreshTokenResponse `json:"refresh_token"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}
