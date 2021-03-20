package interfaces

import (
	"go_jwt_auth/api/models"
	"net/http"
)

type UserRepository interface {
	Save(models.User) (models.User, error)
	FindAll() ([]models.User, error)
	FindById(uint32) (models.User, error)
	Update(uint32, models.UserUpdate) (int64, error)
	// Delete(uint32) (int64, error)
	FindByEmail(string) (models.User, error)
}

type AuthController interface {
	Login(w http.ResponseWriter, r *http.Request)
	// Register(w http.ResponseWriter, r *http.Request)
}
