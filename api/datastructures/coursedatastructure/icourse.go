package coursedatastructure

import (
	"go_jwt_auth/api/models"
	"net/http"
)

type ICourseRepository interface {
	BeginTx() error
	RollbackTx() error
	CommitTx() error

	FindAll() ([]models.Course, error)
	Save(models.Course) (models.Course, error)

	SaveVideoLecture(models.VideoLecture) (models.VideoLecture, error)
	TxSaveVideoLecture(lecture models.VideoLecture) error

	StudentPurchaseCourse(models.StudentCourse) (*models.StudentCourse, error)
	GetCoursesByUserID(uint32) ([]models.Course, error)
	GetVideoLecturesByCourseID(uint32) ([]models.VideoLecture, error)
}

type ICourseController interface {
	CreateCourse(w http.ResponseWriter, r *http.Request)
	GetAllCourses(w http.ResponseWriter, r *http.Request)

	CreateLecture(w http.ResponseWriter, r *http.Request)

	StudentPurchaseCourse(w http.ResponseWriter, r *http.Request)
	GetCoursesByUserID(w http.ResponseWriter, r *http.Request)

	GetVideoLecturesByCourseID(w http.ResponseWriter, r *http.Request)
}
