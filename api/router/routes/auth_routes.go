package routes

import (
	"go_jwt_auth/api/datastructures/userdatastructure"
	"net/http"
)

type authRoutesImpl struct {
	authController userdatastructure.IAuthController
}

func NewAuthRoutes(authController userdatastructure.IAuthController) authRoutesImpl {
	return authRoutesImpl{authController}
}

func (r *authRoutesImpl) Routes() []Route {
	return []Route{
		{
			Uri:     "/auth/login",
			Method:  http.MethodPost,
			Handler: r.authController.Login,
		},
		{
			Uri:     "/auth/register",
			Method:  http.MethodPost,
			Handler: r.authController.UserRegister,
		},
	}
}
