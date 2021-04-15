package courserepository

import (
	"database/sql"
	"go_jwt_auth/api/models"
	"go_jwt_auth/api/utils/channels"
	"go_jwt_auth/api/utils/queryutil"
	"time"
)

type courseRepositoryImpl struct {
	db *sql.DB
	tx *sql.Tx
}

func NewCourseRepository(db *sql.DB) *courseRepositoryImpl {
	return &courseRepositoryImpl{db, nil}
}

func (r *courseRepositoryImpl) BeginTx() error {
	var err error

	r.tx, err = r.db.Begin()
	return err
}

func (r *courseRepositoryImpl) RollbackTx() error {

	err := r.tx.Rollback()
	return err
}

func (r *courseRepositoryImpl) CommitTx() error {
	err := r.tx.Commit()
	return err
}

func (r *courseRepositoryImpl) FindAll() ([]models.Course, error) {
	var err error
	var rows *sql.Rows

	courses := []models.Course{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)

		rows, err = r.db.Query(queryutil.GET_ALL_COURSES)
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)

	if channels.OK(done) {
		for rows.Next() {
			course := models.Course{}
			err = rows.Scan(&course.ID, &course.AuthorID, &course.Title, &course.Description, &course.Price, &course.ImageURL, &course.CreatedAt, &course.UpdatedAt, &course.DeletedAt)

			if err != nil {
				return nil, err
			}

			courses = append(courses, course)
		}
		err = rows.Err() // get any error encountered in iteration
		if err != nil {
			return nil, err
		}

		return courses, nil
	}
	return nil, err

}

func (r *courseRepositoryImpl) Save(course models.Course) (models.Course, error) {
	var err error
	done := make(chan bool)
	go func(ch chan<- bool) {
		_, err = r.db.Exec(queryutil.INSERT_NEW_COURSE, course.Title, course.Description, course.Price, course.ImageURL, time.Now(), time.Now(), course.AuthorID)
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)

	if channels.OK(done) {
		return course, nil
	}
	return models.Course{}, err
}

func (r *courseRepositoryImpl) TxSaveVideoLecture(lecture models.VideoLecture) error {
	var err error

	done := make(chan bool)
	go func(ch chan<- bool) {
		_, err = r.tx.Exec(queryutil.INSERT_NEW_VIDEO_LECTURE, lecture.Title, lecture.VideoURL, time.Now(), time.Now(), lecture.CourseID)
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)

	if channels.OK(done) {
		return nil
	}
	return err
}

func (r *courseRepositoryImpl) SaveVideoLecture(lecture models.VideoLecture) (models.VideoLecture, error) {
	var err error
	done := make(chan bool)
	go func(ch chan<- bool) {
		_, err = r.db.Exec(queryutil.INSERT_NEW_VIDEO_LECTURE, lecture.Title, lecture.VideoURL, time.Now(), time.Now(), lecture.CourseID)
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)

	if channels.OK(done) {
		return lecture, nil
	}
	return models.VideoLecture{}, err
}

func (r *courseRepositoryImpl) StudentPurchaseCourse(sc models.StudentCourse) (*models.StudentCourse, error) {
	var err error
	done := make(chan bool)
	go func(ch chan<- bool) {
		_, err = r.db.Exec(queryutil.INSERT_STUDENT_COURSE, sc.StudentID, sc.CourseID)
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)

	if channels.OK(done) {
		return &sc, nil
	}
	return nil, err
}

func (r *courseRepositoryImpl) GetCoursesByUserID(sID uint32) ([]models.Course, error) {
	var err error
	var rows *sql.Rows
	courses := []models.Course{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)

		rows, err = r.db.Query(queryutil.GET_ALL_COURSES_AGGR_STUDENT_COURSE_AGGR_STUDENT_BY_USER_ID, sID)
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)

	if channels.OK(done) {
		for rows.Next() {
			course := models.Course{}
			err = rows.Scan(
				&course.ID,
				&course.AuthorID,
				&course.Title,
				&course.Description,
				&course.Price,
				&course.ImageURL,
				&course.CreatedAt,
				&course.UpdatedAt,
				&course.DeletedAt,
			)

			if err != nil {
				return nil, err
			}

			courses = append(courses, course)
		}
		err = rows.Err() // get any error encountered in iteration
		if err != nil {
			return nil, err
		}

		return courses, nil
	}
	return nil, err
}

func (r *courseRepositoryImpl) GetVideoLecturesByCourseID(cID uint32) ([]models.VideoLecture, error) {
	var err error
	var rows *sql.Rows
	vls := []models.VideoLecture{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)

		rows, err = r.db.Query(queryutil.GET_ALL_VIDEO_LECTURES_BY_COURSE_ID, cID)
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)

	if channels.OK(done) {
		for rows.Next() {
			vl := models.VideoLecture{}
			err = rows.Scan(
				&vl.ID,
				&vl.Title,
				&vl.VideoURL,
				&vl.CreatedAt,
				&vl.UpdatedAt,
				&vl.DeletedAt,
				&vl.CourseID,
			)

			if err != nil {
				return nil, err
			}

			vls = append(vls, vl)
		}
		err = rows.Err() // get any error encountered in iteration
		if err != nil {
			return nil, err
		}

		return vls, nil
	}
	return nil, err
}
