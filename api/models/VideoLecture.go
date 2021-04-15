package models

import (
	"encoding/json"
	"io"
	"time"
)

type VideoLecture struct {
	ID        uint32     `gorm:"primary_key;auto_increment" json:"id"`
	CourseID  uint32     `json:"course_id"`
	Title     string     `json:"title"`
	VideoURL  string     `json:"video_url"`
	CreatedAt *time.Time `gorm:"default:current_timestamp()" json:"created_at"`
	UpdatedAt *time.Time `gorm:"default:current_timestamp()" json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func (l *VideoLecture) FromJSON(r io.Reader) error {
	json.NewDecoder(r).Decode(l)
	return nil
}
