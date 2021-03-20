package routes

import (
	"go_jwt_auth/api/interfaces"
	"net/http"
)

type authRoutesImpl struct {
	authController interfaces.AuthController
}

func NewAuthRoutes(authController interfaces.AuthController) authRoutesImpl {
	return authRoutesImpl{authController}
}

func (r *authRoutesImpl) Routes() []Route {
	return []Route{
		{
			Uri:     "/auth/login",
			Method:  http.MethodPost,
			Handler: r.authController.Login,
		},
	}
}
