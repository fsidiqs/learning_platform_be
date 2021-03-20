package coursedatastructure

import (
	"encoding/json"
	"io"
)

type VideoLecture struct {
	ID       uint32 `json:"id"`
	CourseID uint32 `json:"course_id"`
	Title    string `json:"title"`
	VideoURL string `json:"video_url"`
}

func (l *VideoLecture) FromJSON(r io.Reader) error {
	json.NewDecoder(r).Decode(l)
	return nil
}

type VideoLectureArr []VideoLecture

func (l *VideoLectureArr) FromJSON(r io.Reader) error {
	json.NewDecoder(r).Decode(&l)
	return nil
}
