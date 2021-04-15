package studentds

import "go_jwt_auth/api/models"

type IStudentRepository interface {
	Save(models.Student) (*models.Student, error)

	// FindAll() ([]models.Student, error)
	// FindById(uint32) (models.Student, error)
	FindByUserID(uint32) (*models.Student, error)
	// Deactivate(uint32) (int64, error)
	// Activate(uint32) (int64, error)
}

// type IUserController interface {
// 	GetUsers(w http.ResponseWriter, r *http.Request)
// 	CreateUser(w http.ResponseWriter, r *http.Request)
// 	GetUser(w http.ResponseWriter, r *http.Request)
// 	GetUserByEmail(w http.ResponseWriter, r *http.Request)
// 	UpdateUser(w http.ResponseWriter, r *http.Request)
// 	DeleteUser(w http.ResponseWriter, r *http.Request)
// 	DeactivateUser(w http.ResponseWriter, r *http.Request)
// 	ActivateUser(w http.ResponseWriter, r *http.Request)
// }
