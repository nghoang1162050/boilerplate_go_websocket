package utils

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
    Roles []string `json:"roles"`
    jwt.RegisteredClaims
}

func GenerateJWTToken(username string, roles []string) (string, int64, error) {
	secret := os.Getenv("JWT_SECRET")
	jwt_expired, _ := strconv.Atoi(os.Getenv("JWT_EXPIRATION"))
	expired := time.Now().Add(time.Duration(jwt_expired) * time.Minute).Unix()

	claims := &CustomClaims{
        Roles: roles,
        RegisteredClaims: jwt.RegisteredClaims{
            ID:        GenerateUserID(username),
            Audience:  jwt.ClaimStrings{username},
            Issuer:    username,
            Subject:   username,
            IssuedAt:  jwt.NewNumericDate(time.Now()),
            ExpiresAt: jwt.NewNumericDate(time.Unix(expired, 0)),
        },
    }

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	
	tokenString, err := token.SignedString([]byte(secret))

	return tokenString, expired, err
}

func ExtractUsernameFromToken(header string) (string, error) {
	tokenString, _ := ExtractTokenFromHeader(header)
	secret := os.Getenv("JWT_SECRET")
	claims := &CustomClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil || !token.Valid {
		return "", err
	}

	return claims.Subject, nil
}

func ExtractRolesFromToken(header string) ([]string, error) {
	tokenString, _ := ExtractTokenFromHeader(header)
	secret := os.Getenv("JWT_SECRET")
	claims := &CustomClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	return claims.Roles, nil
}

func ExtractTokenFromHeader(header string) (string, error) {
	parts := strings.Split(header, " ")
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		return "", fmt.Errorf("invalid Authorization header")
	}
	return parts[1], nil
}
