package usercontroller

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"go_jwt_auth/api/datastructures/userdatastructure"
	"go_jwt_auth/api/models"
	"go_jwt_auth/api/utils/responses"
	"go_jwt_auth/api/authentication"


	"github.com/gorilla/mux"
)

type userControllerImpl struct {
	UserRepository userdatastructure.IUserRepository
}

func NewUserController(repo userdatastructure.IUserRepository) *userControllerImpl {
	return &userControllerImpl{repo}
}

func (c *userControllerImpl) DeactivateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	userResp, err := c.UserRepository.Deactivate(uint32(uid))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s", r.Host, r.RequestURI))
	responses.JSON(w, http.StatusOK, userResp)

}

func (c *userControllerImpl) ActivateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	userResp, err := c.UserRepository.Activate(uint32(uid))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s", r.Host, r.RequestURI))
	responses.JSON(w, http.StatusOK, userResp)

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

	user.Password, err = authentication.Hash(user.Password)

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

func (c *userControllerImpl) GetUserByEmail(w http.ResponseWriter, r *http.Request) {
	var request map[string]string

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, errors.New("expected email"))
		return
	}
	userResp, err := c.UserRepository.FindByEmail(request["email"])

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
