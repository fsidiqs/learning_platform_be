package routes

import (
	"go_jwt_auth/api/datastructures/studentcourseds"
	// "net/http"
)

type studentCourseImpl struct {
	scController studentcourseds.IStudentCourseControlller
}

func NewStudentCourseRoutes(scCont studentcourseds.IStudentCourseControlller) studentCourseImpl {
	return studentCourseImpl{scCont}
}

func (r *studentCourseImpl) Routes() []Route {
	return []Route{
		// {
		// 	Uri:     "/students/courses",
		// 	Method:  http.MethodGet,
		// 	Handler: r.scController.GetCoursesByUserID,
		// },
	
	}
}
