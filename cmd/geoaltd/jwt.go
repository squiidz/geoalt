package main

import (
	"context"
	"crypto/sha256"
	"errors"
	"log"

	jwt "github.com/dgrijalva/jwt-go"
	geo "github.com/squiidz/geoalt"
	"google.golang.org/grpc/metadata"
)

const Secret = "somenotsosecretsecret"

type Claim struct {
	ID    uint32 `json:"id"`
	Email string `json:"email"`
}

func (c *Claim) Valid() error {
	if c.Email != "" {
		return nil
	}
	return errors.New("Invalid claim")
}

func hashPassword(p string) string {
	h := sha256.New()
	h.Write([]byte(p))
	return string(h.Sum(nil))
}

func genToken(u *geo.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claim{
		ID:    u.ID,
		Email: u.Email,
	})
	signToken, err := token.SignedString([]byte(Secret))
	if err != nil {
		return "", err
	}
	return signToken, nil
}

func tokenIsValid(t string) (*Claim, bool) {
	var claim Claim
	_, err := jwt.ParseWithClaims(t, &claim, func(token *jwt.Token) (interface{}, error) {
		return []byte(Secret), nil
	})
	if err != nil {
		log.Println(err)
		return &claim, false
	}
	return &claim, true
}

func checkToken(ctx context.Context) (*Claim, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok || len(md.Get("token")) <= 0 {
		return nil, errors.New("No Metadata")
	}
	c, ok := tokenIsValid(md.Get("token")[0])
	if !ok {
		return nil, errors.New("Invalid Token please login")
	}
	return c, nil
}
