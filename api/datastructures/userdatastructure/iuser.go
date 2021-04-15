package userdatastructure

import (
	"go_jwt_auth/api/models"
	"net/http"
)

type IUserRepository interface {
	Save(models.User) (models.User, error)
	Update(uint32, models.UserUpdate) (int64, error)
	Delete(uint32) (int64, error)

	FindAll() ([]models.User, error)
	FindById(uint32) (models.User, error)
	FindByEmail(string) (*models.User, error)
	FindByEmailWithPassword(string) (*models.User, error)

	Deactivate(uint32) (int64, error)
	Activate(uint32) (int64, error)
}

type IUserController interface {
	GetUsers(w http.ResponseWriter, r *http.Request)
	CreateUser(w http.ResponseWriter, r *http.Request)
	GetUser(w http.ResponseWriter, r *http.Request)
	GetUserByEmail(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
	DeactivateUser(w http.ResponseWriter, r *http.Request)
	ActivateUser(w http.ResponseWriter, r *http.Request)
}

type IAuthController interface {
	Login(w http.ResponseWriter, r *http.Request)
	UserRegister(w http.ResponseWriter, r *http.Request)
}
