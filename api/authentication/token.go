package authentication

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type UserPublic struct {
	Email string
}

type Claim struct {
	User UserPublic `json:"user"`
	jwt.StandardClaims
}

func GenerateJWT(user UserPublic, secret []byte) (string, error) {
	claim := Claim{
		User: user,
		StandardClaims: jwt.StandardClaims{
			Issuer:    "distrodakwah.id",
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString(secret)

}
