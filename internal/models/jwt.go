package models

import (
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	ID          int         `json:"id"`
	Name        string      `json:"name"`
	Email       string      `json:"email"`
	Provider    string      `json:"provider"`
	Address     string      `json:"address"`
	Authorities []Authority `json:"authorities"`
	jwt.RegisteredClaims
}

type AddrClaims struct {
	MemberID int    `json:"memberId"`
	Addr     string `json:"address"`
	jwt.RegisteredClaims
}
