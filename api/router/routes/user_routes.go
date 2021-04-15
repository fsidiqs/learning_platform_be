package routes

import (
	"go_jwt_auth/api/datastructures/userdatastructure"
	"net/http"
)

type userRoutesImpl struct {
	userController userdatastructure.IUserController
}

func NewUserRoutes(userController userdatastructure.IUserController) userRoutesImpl {
	return userRoutesImpl{userController}
}

func (r *userRoutesImpl) Routes() []Route {
	return []Route{
		{
			Uri:     "/users",
			Method:  http.MethodGet,
			Handler: r.userController.GetUsers,
		},
		{
			Uri:     "/users",
			Method:  http.MethodPost,
			Handler: r.userController.CreateUser,
		},
		{
			Uri:     "/users/by-email",
			Method:  http.MethodGet,
			Handler: r.userController.GetUserByEmail,
		},
		{
			Uri:     "/users/{id}/deactivate",
			Method:  http.MethodPost,
			Handler: r.userController.DeactivateUser,
		},
		{
			Uri:     "/users/{id}/activate",
			Method:  http.MethodPost,
			Handler: r.userController.ActivateUser,
		},
		{
			Uri:     "/users/{id}",
			Method:  http.MethodGet,
			Handler: r.userController.GetUser,
		},
		{
			Uri:     "/users/{id}",
			Method:  http.MethodPut,
			Handler: r.userController.UpdateUser,
		},
		{
			Uri:     "/users/{id}",
			Method:  http.MethodDelete,
			Handler: r.userController.DeleteUser,
		},
	}
}
