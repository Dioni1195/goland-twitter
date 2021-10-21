package jwt

import (
	"GitHub/goland-twitter/models"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func GenerateToken(u models.User) (string, error) {

	secretKey := []byte("SecretKey")

	payload := jwt.MapClaims{
		"email":    u.Email,
		"name":     u.Name,
		"lastName": u.LastName,
		"bio":      u.Bio,
		"_id":      u.Id.Hex(),
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	tokenStr, err := token.SignedString(secretKey)
	if err != nil {
		return tokenStr, err
	}

	return tokenStr, nil
}
