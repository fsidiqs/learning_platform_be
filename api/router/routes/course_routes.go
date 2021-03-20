package routes

import (
	"go_jwt_auth/api/interfaces"
	"net/http"
)

type courseRoutesImpl struct {
	courseController interfaces.ICourseController
}

func NewCourseRoutes(courseController interfaces.ICourseController) courseRoutesImpl {
	return courseRoutesImpl{courseController}
}

func (r *courseRoutesImpl) Routes() []Route {
	return []Route{
		{
			Uri:     "/courses",
			Method:  http.MethodGet,
			Handler: r.courseController.GetAllCourses,
		},
		{
			Uri:     "/courses",
			Method:  http.MethodPost,
			Handler: r.courseController.CreateCourse,
		},
		// {
		// 	Uri:     "/users/{id}",
		// 	Method:  http.MethodGet,
		// 	Handler: r.courseController.GetUser,
		// },
		// {
		// 	Uri:     "/users/{id}",
		// 	Method:  http.MethodPut,
		// 	Handler: r.courseController.UpdateUser,
		// },
		// {
		// 	Uri:     "/users/{id}",
		// 	Method:  http.MethodDelete,
		// 	Handler: r.courseController.DeleteUser,
		// },
	}
}
