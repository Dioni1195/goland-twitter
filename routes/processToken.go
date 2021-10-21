package routes

import (
	"GitHub/goland-twitter/bd"
	"GitHub/goland-twitter/models"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"strings"
)

var Email string
var UserID string

func ProcessToken(token string) (*models.Claim, bool, string, error) {
	secretKey := []byte("SecretKey")
	claims := &models.Claim{}

	splitToken := strings.Split(token, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, "", errors.New("formato de token invalido")
	}

	token = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(token, claims, func(tk *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err == nil {
		_, found, ID := bd.CheckUserAlreadyExist(claims.Email)
		if found {
			Email = claims.Email
			UserID = ID
		}
		return claims, found, ID, nil
	}

	if !tkn.Valid {
		return claims, false, string(""), errors.New("tpken invalido")
	}
	return claims, false, string(""), err

}
