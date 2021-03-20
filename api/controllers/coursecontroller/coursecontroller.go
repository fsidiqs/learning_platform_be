package coursecontroller

import (
	"errors"
	"fmt"
	coursedatastructure "go_jwt_auth/api/datastructures/coursedatastructre"
	"go_jwt_auth/api/interfaces"
	ifilestorage "go_jwt_auth/api/interfaces/filestorage"
	"strings"

	"go_jwt_auth/api/utils/responses"

	"net/http"
)

type courseControllerImpl struct {
	CourseRepository interfaces.ICourseRepository
	FileStorage      ifilestorage.IFileStorage
}

func NewCourseController(repo interfaces.ICourseRepository, fs ifilestorage.IFileStorage) *courseControllerImpl {
	return &courseControllerImpl{CourseRepository: repo, FileStorage: fs}
}

func (c *courseControllerImpl) GetAllCourses(w http.ResponseWriter, r *http.Request) {
	userResp, err := c.CourseRepository.FindAll()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusCreated, userResp)
}

func (c *courseControllerImpl) CreateCourse(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(128 * 1024)
	if err != nil {
		http.Error(w, "Expected multipart form data", http.StatusBadRequest)
		return
	}

	var courseReq coursedatastructure.CourseCreateReq
	var videoLecArr coursedatastructure.VideoLectureArr
	// vide coursedatastructure.VideoLectureArr{}

	courseForm := r.FormValue("course")
	err = courseReq.FromJSON(strings.NewReader(courseForm))
	err = videoLecArr.FromJSON(strings.NewReader(courseReq.VideoLectures))
	fmt.Printf("%+v \n", videoLecArr)
	return
	course := courseReq.ToModel()
	course.AuthorID = 1

	// err =
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, errors.New("invalid course datas"))
		return
	}

	// ff, mh, err := r.FormFile("course_thumbnail")
	_, _, err = r.FormFile("course_thumbnail")

	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, errors.New("expected Course Thumbnail file"))
		return
	}

	// res, err := c.FileStorage.UploadImage("/test.jpg", ff)
	res := "file.jpg"
	course.ImageURL = res
	courseResp, err := c.CourseRepository.Save(course)

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusCreated, courseResp)

}
