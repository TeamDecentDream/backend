package models

import (
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	Name        string      `json:"name"`
	Email       string      `json:"email"`
	Provider    string      `json:"provider"`
	Authorities []Authority `json:"authorities"`
	jwt.RegisteredClaims
}
