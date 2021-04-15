package models

import (
	"encoding/json"
	"html"
	"io"
	"strings"
	"time"
)

// CONST (
// 	ERR_ = ""
// )

type User struct {
	ID        uint32     `gorm:"primary_key;auto_increment" json:"id"`
	RoleID    uint8      `json:"role_id"`
	Name      string     `json:"name"`
	Email     string     `gorm:"size:50;not null; unique" json:"email"`
	Password  string     `json:"password"`
	CreatedAt *time.Time `gorm:"default:current_timestamp()" json:"created_at"`
	UpdatedAt *time.Time `gorm:"default:current_timestamp()" json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func (u *User) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	e.Decode(u)
	return nil
}

type UserUpdate struct {
	Name        string     `json:"name"`
	Email       string     `gorm:"size:50;not null; unique" json:"email"`
	UpdatedAt   *time.Time `gorm:"default:current_timestamp()" json:"updated_at"`
	UpdateField string     `json:"update_field"`
}

func (u *UserUpdate) Prepare() {
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))

}

func (u *UserUpdate) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	e.Decode(u)
	return nil
}

type UserWStudent struct {
	ID        uint32     `gorm:"primary_key;auto_increment" json:"id"`
	RoleID    uint8      `json:"role_id"`
	Name      string     `json:"name"`
	Email     string     `gorm:"size:50;not null; unique" json:"email"`
	Password  string     `json:"password"`
	CreatedAt *time.Time `gorm:"default:current_timestamp()" json:"created_at"`
	UpdatedAt *time.Time `gorm:"default:current_timestamp()" json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`

	Student Student `json:"student"`
}
