package authentication

import (
	"fmt"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	jwtReq "github.com/dgrijalva/jwt-go/request"
)

type UserJWTClaim struct {
	Email string
}

type Claim struct {
	User UserJWTClaim `json:"user"`
	jwt.StandardClaims
}

func GenerateJWT(user UserJWTClaim, secret []byte) (string, error) {
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

func PerformAuthChekingFromReq(r *http.Request, jwtSecret []byte) (*jwt.Token, error) {
	token, err := jwtReq.ParseFromRequestWithClaims(
		r,
		jwtReq.OAuth2Extractor,
		&Claim{},
		func(t *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		},
	)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return token, err
}
