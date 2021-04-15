package coursedatastructure

import (
	"encoding/json"
	"errors"
	"fmt"
	"go_jwt_auth/api/models"
	"io"
	"net/textproto"
	"strconv"
	"strings"
	"time"
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

func (l *VideoLecture) ToModel() models.VideoLecture {
	return models.VideoLecture{
		CourseID: l.CourseID,
		Title:    l.Title,
		VideoURL: l.VideoURL,
	}
}

type VideoLectureArr []VideoLecture

func (vArr *VideoLectureArr) FromJSON(r io.Reader) error {
	json.NewDecoder(r).Decode(&vArr)
	return nil
}

func (vArr *VideoLectureArr) FillCourseID(id uint64) {
	for i := 0; i < len(*vArr); i++ {
		(*vArr)[i].CourseID = uint32(id)
	}
}

func (vArr *VideoLectureArr) FillVideoLectureLoc(fileLocs []LectureFile) {
	vLen := len(*vArr)
	for i := 0; i < vLen; i++ {
		// find video location
		for _, v := range fileLocs {
			// convert key to int
			keyInt, err := strconv.Atoi(v.Key)
			if err != nil {
				fmt.Println("error parsing key")
			}

			if keyInt == i {

				(*vArr)[i].VideoURL = v.Location
			}
		}

	}

}

func (vArr *VideoLectureArr) ToModel() []models.VideoLecture {
	mVArr := make([]models.VideoLecture, len(*vArr))
	for i, v := range *vArr {
		now := time.Now()
		mVArr[i] = models.VideoLecture{
			CourseID:  v.CourseID,
			Title:     v.Title,
			VideoURL:  v.VideoURL,
			CreatedAt: &now,
			UpdatedAt: &now,
		}
	}

	return mVArr
}

type LectureFile struct {
	Location string
	Key      string
}

func (l *LectureFile) FillWithFileHeader(h textproto.MIMEHeader, uploadLoc string) error {
	fh := h.Get("Content-Disposition")

	// get key of formfile
	split := strings.Split(fh, " ")
	formName := strings.Split(split[1], `"`)
	if len(formName) < 3 {
		return errors.New("file header must not be empty")
	}
	l.Key = formName[1]
	l.Location = uploadLoc
	return nil
}
