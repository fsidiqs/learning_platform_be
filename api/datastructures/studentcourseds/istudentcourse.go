package studentcourseds

import (
	"go_jwt_auth/api/models"
	"net/http"
)

type IStudentCourseRepository interface {
	GetCoursesByUserID(uint32) ([]models.Course, error)
}

type IStudentCourseControlller interface {
	GetCoursesByUserID(w http.ResponseWriter, r *http.Request)
}