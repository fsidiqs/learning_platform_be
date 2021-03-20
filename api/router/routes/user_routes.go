package routes

import (
	"go_jwt_auth/api/controllers/usercontroller"
	"net/http"
)

type userRoutesImpl struct {
	userController usercontroller.UserController
}

func NewUserRoutes(userController usercontroller.UserController) userRoutesImpl {
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
