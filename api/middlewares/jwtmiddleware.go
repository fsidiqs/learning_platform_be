package middlewares

import (
	"context"
	"errors"
	"net/http"

	"go_jwt_auth/api/authentication"
	"go_jwt_auth/api/datastructures/userdatastructure"
	"go_jwt_auth/api/utils/contextkey"
	"go_jwt_auth/api/utils/responses"
)

type AuthConfig struct {
	JwtSecret []byte
	UserRepo  userdatastructure.IUserRepository
}

var SetMiddlewareAuth = func(authconfig AuthConfig) Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {

		return func(w http.ResponseWriter, r *http.Request) {

			token, err := authentication.PerformAuthChekingFromReq(r, authconfig.JwtSecret)
			if err != nil {
				responses.ERROR(w, http.StatusNotFound, errors.New("Unauthenticated"))
				return
			}

			user, err := authconfig.UserRepo.FindByEmail(token.Claims.(*authentication.Claim).User.Email)
			if err != nil {
				responses.ERROR(w, http.StatusNotFound, errors.New("authentication failed"))
				return
			}
			userContext := context.WithValue(
				r.Context(),
				contextkey.UserKey("user"),
				user,
			)

			next(w, r.WithContext(userContext))

		}

	}

}

// var SetAdminMiddlewareAuth = func()
