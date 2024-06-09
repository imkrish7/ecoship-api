package services

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/cristalhq/jwt/v5"
)

type UserClaims struct {
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.RegisteredClaims
}

var secret = []byte("secret")

func GenerateToken(email string, role string) (string, error) {

	claims := &UserClaims{
		Email: email,
		Role:  role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			Issuer:    "ecoship",
		},
	}
	signer, _ := jwt.NewSignerHS(jwt.HS256, secret)
	unsigned := jwt.NewBuilder(signer)

	token, err := unsigned.Build(claims)

	if err != nil {
		return "", err
	}
	return token.String(), nil
}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	verifier, _ := jwt.NewVerifierHS(jwt.HS256, secret)
	token, err := jwt.Parse([]byte(tokenString), verifier)

	if err != nil {
		return nil, err
	}

	var newClaims UserClaims
	json.Unmarshal(token.Claims(), &newClaims)
	valid := newClaims.IsValidAt(time.Now().UTC())
	if !valid {
		return nil, errors.New("Token has been expired")
	}
	return token, nil
}
