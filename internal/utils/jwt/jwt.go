package jwt

import (
	"backend/internal/models"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var (
	secretKey = []byte("likelionbcs4nextfarmdefiproject")
)

func AccessTokenProvider(member *models.Member) (string, error) {
	var claims models.Claims
	claims.ID = member.ID
	claims.Name = member.Name
	claims.Email = member.Email
	claims.Authorities = member.Authorities
	claims.RegisteredClaims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(2 * time.Hour))

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	accessToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return accessToken, err
}

func AccessTokenVerifier(accessToken string) (int, string, string, []models.Authority, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})
	if err != nil {
		return -1, "", "", nil, err
	}
	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
	} else {
		return -1, "", "", nil, fmt.Errorf("invalid token")
	}

	parseResult, err := jwt.ParseWithClaims(accessToken, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("likelionbcs4nextfarmdefiproject"), nil
	})
	if err != nil {
		return -1, "", "", nil, err
	}

	if claims, ok := parseResult.Claims.(*models.Claims); ok && parseResult.Valid {
		return claims.ID, claims.Email, claims.Name, claims.Authorities, nil
	} else {
		return -1, "", "", nil, err
	}
}

func AddressTokenProvider(id int, addr string) (string, error) {
	var claims models.AddrClaims
	claims.MemberID = id
	claims.Addr = addr
	claims.RegisteredClaims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(15 * time.Second))

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	accessToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return accessToken, err
}
