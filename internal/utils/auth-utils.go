package utils

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func HashedPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateUserID(username string) string {
    id := uuid.NewSHA1(uuid.NameSpaceDNS, []byte(username))
    return id.String()
}

func ShouldIgnoreRequest(path string) bool {
	allowing := map[string]bool{
		"/api/auth/login":    true,
		"/api/auth/register": true,
		"/swagger/*":          true,
		"/metrics":          true,
	}

    return allowing[path]
}
