package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const SECRET_KEY = "DUMMY_KEY"

func GenerateJWT(email string, userID int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userID": userID,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(SECRET_KEY))
}

func ValidateJWT(token string) error {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("Unexpected signing error")
		}

		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		return errors.New("could not parse token")
	}

	if !parsedToken.Valid {
		return errors.New("Invalid token")
	}

	_, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
		return errors.New("invalid token claims")
	}

	//email := claims["email"].(string)
	//userID := claims["userID"].(int64)
	return nil
}
