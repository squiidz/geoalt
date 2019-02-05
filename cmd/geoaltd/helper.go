package main

import (
	"crypto/sha256"

	jwt "github.com/dgrijalva/jwt-go"
	geo "github.com/squiidz/geoalt"
)

const Secret = "somenotsosecretsecret"

func hashPassword(p string) string {
	h := sha256.New()
	return string(h.Sum([]byte(p)))
}

func genToken(u *geo.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": u.Email,
	})
	signToken, err := token.SignedString([]byte(Secret))
	if err != nil {
		return "", err
	}
	return signToken, nil
}

func tokenIsValid(t string) bool {
	// _, err := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
	// 	return token, nil
	// })
	// if err != nil {
	// 	log.Println(err)
	// 	return false
	// }
	return true
}
