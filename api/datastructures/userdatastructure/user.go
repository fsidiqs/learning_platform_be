package userdatastructure

import (
	"encoding/json"
	"go_jwt_auth/api/models"
	"html"
	"io"
	"strings"
)

type UserLoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *UserLoginInput) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	e.Decode(u)
	return nil
}

func (u *UserLoginInput) Prepare() {
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
}

type UserRegisterInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	// RoleID   uint8  `json:"role_id"`
}

func (u *UserRegisterInput) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	e.Decode(u)
	return nil
}

func (u *UserRegisterInput) Prepare() {
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.Name = html.EscapeString(strings.TrimSpace(u.Name))
}

func (u UserRegisterInput) ToModel() models.User {
	return models.User{
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
		RoleID:   2,
	}
}
