package coursedatastructure

import (
	"encoding/json"
	"go_jwt_auth/api/models"
	"io"
)

type CourseCreateReq struct {
	Title         string `json:"title"`
	Description   string `json:"description"`
	Price         uint64 `json:"price"`
	VideoLectures string `json:"video_lectures"`
	//
}

func (c *CourseCreateReq) FromJSON(r io.Reader) error {
	json.NewDecoder(r).Decode(c)
	return nil
}

func (c *CourseCreateReq) ToModel() models.Course {
	return models.Course{
		Title:       c.Title,
		Description: c.Description,
		Price:       c.Price,
	}
}

type CourseImageReq struct {
	Filename  string
	Content   io.Reader
	Size      int64
	URLResult string
}
