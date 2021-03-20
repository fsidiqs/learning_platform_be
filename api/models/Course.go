package models

import (
	"encoding/json"
	"io"
	"time"
)

type Course struct {
	ID            uint32         `gorm:"primary_key;auto_increment" json:"id"`
	AuthorID      uint32         `json:"author_id"`
	Title         string         `json:"title"`
	Description   string         `json:"description"`
	Price         uint64         `json:"price"`
	ImageURL      string         `json:"image_url"`
	CreatedAt     time.Time      `gorm:"default:current_timestamp()" json:"created_at"`
	UpdatedAt     time.Time      `gorm:"default:current_timestamp()" json:"updated_at"`
	DeletedAt     time.Time      `json:"deleted_at"`
	VideoLectures []VideoLecture `gorm:"foreignKey:CourseID;references:ID" json:"video_lectures"`
}

func (c *Course) FromJSON(r io.Reader) error {
	json.NewDecoder(r).Decode(c)
	return nil
}
