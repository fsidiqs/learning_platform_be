package userdatastructure

import (
	"encoding/json"
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
