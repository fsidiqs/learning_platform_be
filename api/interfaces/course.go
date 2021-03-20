package interfaces

import (
	"go_jwt_auth/api/models"
	"net/http"
)

type ICourseRepository interface {
	FindAll() ([]models.Course, error)
	Save(models.Course) (models.Course, error)

	// TxSave(*gorm.DB, models.Course) (*gorm.DB, models.Course, error)
}

type ICourseController interface {
	CreateCourse(w http.ResponseWriter, r *http.Request)
	GetAllCourses(w http.ResponseWriter, r *http.Request)
}
