package coursecontroller

import (
	"errors"
	"fmt"
	"go_jwt_auth/api/datastructures/coursedatastructure"
	"go_jwt_auth/api/datastructures/studentds"
	"go_jwt_auth/api/models"

	ifilestorage "go_jwt_auth/api/datastructures/filestorage"
	"strconv"
	"strings"

	"go_jwt_auth/api/utils/contextkey"
	"go_jwt_auth/api/utils/responses"

	"net/http"

	"github.com/gorilla/mux"
)

type courseControllerImpl struct {
	CourseR     coursedatastructure.ICourseRepository
	StudentR    studentds.IStudentRepository
	FileStorage ifilestorage.IFileStorage
}

func NewCourseController(
	repo coursedatastructure.ICourseRepository,
	studentR studentds.IStudentRepository,
	fs ifilestorage.IFileStorage,
) *courseControllerImpl {
	return &courseControllerImpl{
		CourseR:     repo,
		StudentR:    studentR,
		FileStorage: fs,
	}
}

func (c *courseControllerImpl) GetAllCourses(w http.ResponseWriter, r *http.Request) {
	courseresp, err := c.CourseR.FindAll()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, courseresp)
}

func (c *courseControllerImpl) CreateCourse(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(128 * 1024)
	if err != nil {
		http.Error(w, "Expected multipart form data", http.StatusBadRequest)
		return
	}

	var courseReq coursedatastructure.CourseCreateReq

	courseForm := r.FormValue("course")
	err = courseReq.FromJSON(strings.NewReader(courseForm))

	course := courseReq.ToModel()
	course.AuthorID = 1

	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, errors.New("invalid course datas"))
		return
	}

	ff, _, err := r.FormFile("course_thumbnail")
	// dummy thumbnail
	// _, _, err = r.FormFile("course_thumbnail")

	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, errors.New("expected Course Thumbnail file"))
		return
	}

	res, err := c.FileStorage.UploadImage("/test.jpg", ff)
	// res := "file.jpg"
	course.ImageURL = res
	courseResp, err := c.CourseR.Save(course)

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusCreated, courseResp)

}

func (c *courseControllerImpl) CreateLecture(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	courseID, err := strconv.ParseUint(vars["course_id"], 10, 32)

	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	_ = r.ParseMultipartForm(128 * 1024)

	var videoLecArr coursedatastructure.VideoLectureArr
	lectureForm := r.FormValue("lectures")

	err = videoLecArr.FromJSON(strings.NewReader(lectureForm))

	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	formData := r.MultipartForm

	var lectureFiles []coursedatastructure.LectureFile


	for _, files := range formData.File {
		
		for _, file := range files {
			fCont, err := file.Open()
			
			if err != nil {
				responses.ERROR(w, http.StatusBadRequest, err)

				return
			}
			defer fCont.Close()
			// TODO upload using service
			fmt.Println("uploading video file")
			res, err := c.FileStorage.UploadVideo(file.Filename, fCont)
			// res := "file.mp4"
			fmt.Println(res)
			if err != nil {
				// fmt.Printf("%+v \n", )
				fmt.Println("Error uploading file \n", err)
				return
			}

			// create new lecturefile and fill with current header and upload location
			var newLectureFile coursedatastructure.LectureFile
			err = newLectureFile.FillWithFileHeader(file.Header, res)
			if err != nil {
				fmt.Println("Error naming file \n", err)
				return
			}
			// append to lecturefiles
			lectureFiles = append(lectureFiles, newLectureFile)

		}
	}

	videoLecArr.FillCourseID(courseID)
	videoLecArr.FillVideoLectureLoc(lectureFiles)
	vLectureArr := videoLecArr.ToModel()

	// fill video url
	err = c.CourseR.BeginTx()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)

		return
	}
	vLectLen := len(vLectureArr)
	for i := 0; i < vLectLen; i++ {
		// find video location

		err := c.CourseR.TxSaveVideoLecture(vLectureArr[i])
		if err != nil {
			c.CourseR.RollbackTx()
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			fmt.Println(err)
			return
		}
	}

	err = c.CourseR.CommitTx()
	if err != nil {
		fmt.Println("failed to create lecture")
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
}

// TODO make theirown controlller
func (c *courseControllerImpl) StudentPurchaseCourse(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	courseID, err := strconv.ParseUint(vars["course_id"], 10, 32)

	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	user := &models.User{}
	user = r.Context().Value(contextkey.UserKey("user")).(*models.User)
	student, err := c.StudentR.FindByUserID(user.ID)
	studentCourse := models.StudentCourse{
		StudentID: student.ID,
		CourseID:  uint32(courseID),
	}
	resp, err := c.CourseR.StudentPurchaseCourse(studentCourse)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	responses.JSON(w, http.StatusCreated, resp)
}

func (c *courseControllerImpl) GetCoursesByUserID(w http.ResponseWriter, r *http.Request) {

	user := &models.User{}
	user = r.Context().Value(contextkey.UserKey("user")).(*models.User)

	courses, err := c.CourseR.GetCoursesByUserID(user.ID)
	fmt.Println(user)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, courses)
}

func (c *courseControllerImpl) GetVideoLecturesByCourseID(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	courseID, err := strconv.ParseUint(vars["course_id"], 10, 32)
	lectures, err := c.CourseR.GetVideoLecturesByCourseID(uint32(courseID))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, lectures)
}
