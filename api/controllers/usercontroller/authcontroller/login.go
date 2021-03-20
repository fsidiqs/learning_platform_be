package authcontroller

import (
	"go_jwt_auth/api/datastructures/userdatastructure"
	"go_jwt_auth/api/interfaces"
	"go_jwt_auth/config"

	"go_jwt_auth/api/authentication"
	"go_jwt_auth/api/utils/responses"
	"net/http"
)

type authControllerImpl struct {
	JwtConfig      config.JwtConf
	UserRepository interfaces.UserRepository
}

func NewAuthController(jwtConfig config.JwtConf, repo interfaces.UserRepository) *authControllerImpl {
	return &authControllerImpl{JwtConfig: jwtConfig, UserRepository: repo}
}

func (c *authControllerImpl) Login(w http.ResponseWriter, r *http.Request) {
	userInput := userdatastructure.UserLoginInput{}
	err := userInput.FromJSON(r.Body)

	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
	}
	// TODO validate login data
	// refetch user by find by email
	user, err := c.UserRepository.FindByEmail(userInput.Email)

	err = authentication.VerifyPassword(user.Password, userInput.Password)

	if err != nil {

		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	token, err := authentication.GenerateJWT(authentication.UserPublic{Email: user.Email}, c.JwtConfig.JWTSecret)
	if err != nil {

		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	responses.JSON(w, http.StatusOK, token)

}
