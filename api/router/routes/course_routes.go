package routes

import (
	"go_jwt_auth/api/datastructures/coursedatastructure"

	"net/http"
)

type courseRoutesImpl struct {
	courseController coursedatastructure.ICourseController
}

func NewCourseRoutes(courseController coursedatastructure.ICourseController) courseRoutesImpl {
	return courseRoutesImpl{courseController}
}

func (r *courseRoutesImpl) Routes() []Route {
	return []Route{
		{
			Uri:     "/courses",
			Method:  http.MethodGet,
			Handler: r.courseController.GetAllCourses,
			Auth: Auth{
				AdminRequired: true,
			},
		},
		{
			Uri:     "/courses",
			Method:  http.MethodPost,
			Handler: r.courseController.CreateCourse,
		},
		{
			Uri:     "/courses/{course_id}",
			Method:  http.MethodPost,
			Handler: r.courseController.CreateLecture,
		},
		{
			Uri:     "/courses/{course_id}/purchase",
			Method:  http.MethodPost,
			Handler: r.courseController.StudentPurchaseCourse,
			Auth: Auth{
				AdminRequired: true,
			},
		},
		{
			Uri:     "/courses/students",
			Method:  http.MethodGet,
			Handler: r.courseController.GetCoursesByUserID,
			Auth: Auth{
				AdminRequired: true,
			},
		},
		{
			Uri:     "/courses/{course_id}/video-lectures",
			Method:  http.MethodGet,
			Handler: r.courseController.GetVideoLecturesByCourseID,
			Auth: Auth{
				AdminRequired: true,
			},
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
