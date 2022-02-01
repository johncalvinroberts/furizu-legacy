// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type EmptyResponse struct {
	Success bool `json:"success"`
}

type JwtResponse struct {
	Success      bool   `json:"success"`
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type User struct {
	ID           string `json:"id"`
	Email        string `json:"email"`
	CreatedAt    string `json:"createdAt"`
	LastUpsertAt string `json:"lastUpsertAt"`
}
