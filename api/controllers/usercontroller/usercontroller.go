package usercontroller

import (
	"encoding/json"
	"fmt"
	"go_jwt_auth/api/interfaces"
	"go_jwt_auth/api/models"
	"go_jwt_auth/api/security"
	"go_jwt_auth/api/utils/responses"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type UserController interface {
	GetUsers(w http.ResponseWriter, r *http.Request)
	CreateUser(w http.ResponseWriter, r *http.Request)
	GetUser(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
}

type userControllerImpl struct {
	UserRepository interfaces.UserRepository
}

func NewUserController(repo interfaces.UserRepository) *userControllerImpl {
	return &userControllerImpl{repo}
}

func (c *userControllerImpl) GetUsers(w http.ResponseWriter, r *http.Request) {
	userResp, err := c.UserRepository.FindAll()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusCreated, userResp)
}

func (c *userControllerImpl) CreateUser(w http.ResponseWriter, r *http.Request) {

	user := models.User{}
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	user.Password, err = security.Hash(user.Password)

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	userResp, err := c.UserRepository.Save(user)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, userResp.ID))
	responses.JSON(w, http.StatusCreated, userResp)

}

func (c *userControllerImpl) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 32)

	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	userResp, err := c.UserRepository.FindById(uint32(uid))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s", r.Host, r.RequestURI))
	responses.JSON(w, http.StatusOK, userResp)
}

func (c *userControllerImpl) UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	user := models.UserUpdate{}
	err = user.FromJSON(r.Body)

	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	userResp, err := c.UserRepository.Update(uint32(uid), user)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s", r.Host, r.RequestURI))
	responses.JSON(w, http.StatusCreated, userResp)
}

func (c *userControllerImpl) DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete user"))
}
