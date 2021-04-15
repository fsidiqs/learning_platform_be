package authcontroller

import (
	"errors"
	"go_jwt_auth/api/datastructures/studentds"
	"go_jwt_auth/api/datastructures/userdatastructure"
	"go_jwt_auth/api/models"
	"go_jwt_auth/config"

	"go_jwt_auth/api/authentication"
	"go_jwt_auth/api/utils/responses"
	"net/http"
)

type authControllerImpl struct {
	JwtConfig config.JwtConf
	UserR     userdatastructure.IUserRepository
	StudentR  studentds.IStudentRepository
}

func NewAuthController(jwtConfig config.JwtConf, userR userdatastructure.IUserRepository, studentR studentds.IStudentRepository) *authControllerImpl {
	return &authControllerImpl{JwtConfig: jwtConfig, UserR: userR, StudentR: studentR}
}

func (c *authControllerImpl) Login(w http.ResponseWriter, r *http.Request) {
	userInput := userdatastructure.UserLoginInput{}
	err := userInput.FromJSON(r.Body)

	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
	}
	// TODO validate login data
	// refetch user by find by email
	user, err := c.UserR.FindByEmailWithPassword(userInput.Email)
	if user == nil {
		responses.ERROR(w, http.StatusBadRequest, errors.New("wrong credentials"))
		return
	}

	err = authentication.VerifyPassword(user.Password, userInput.Password)

	if err != nil {

		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	token, err := authentication.GenerateJWT(authentication.UserJWTClaim{Email: user.Email}, c.JwtConfig.JWTSecret)
	if err != nil {

		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	responses.JSON(w, http.StatusOK, token)

}

func (c *authControllerImpl) UserRegister(w http.ResponseWriter, r *http.Request) {

	var regUserInput userdatastructure.UserRegisterInput
	err := regUserInput.FromJSON(r.Body)

	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
	}

	regUserInput.Password, err = authentication.Hash(regUserInput.Password)

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	user := regUserInput.ToModel()
	userResp, err := c.UserR.Save(user)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	userResp.Password = ""
	studentResp, err := c.StudentR.Save(models.Student{UserID: userResp.ID})
	_ = studentResp

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusCreated, userResp)

}
