package auth

import (
	"fmt"
	"log"

	"github.com/golang-jwt/jwt/v5"
)

var secret = []byte("jajcocjaj")

type TokenClaims struct {
	Id   int64
	Name string
}

func ValidateToken(tokenString *string) (*jwt.MapClaims, error) {
	if len(*tokenString) == 0 {
		return nil, fmt.Errorf("empty token")
	}
	token, err := jwt.Parse(*tokenString, func(token *jwt.Token) (interface{}, error) {
		if alg, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		} else {
			if alg == nil {
				return nil, fmt.Errorf("unexpected alg %s", token.Header["alg"])
			}
		}

		return secret, nil
	})
	if err != nil {
		log.Println(*tokenString)
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		return nil, fmt.Errorf("error parsing token claims")
	}

	return &claims, nil
}
