package studentrepository

import (
	"database/sql"
	"go_jwt_auth/api/models"
	"go_jwt_auth/api/utils/channels"
	"go_jwt_auth/api/utils/queryutil"
)

type studentRepositoryImpl struct {
	db *sql.DB
}

func NewStudentRepository(db *sql.DB) *studentRepositoryImpl {
	return &studentRepositoryImpl{db}
}

func (r *studentRepositoryImpl) Save(student models.Student) (*models.Student, error) {
	var err error
	done := make(chan bool)
	go func(ch chan<- bool, student models.Student) {
		_, err = r.db.Exec(queryutil.INSERT_NEW_STUDENT, student.UserID)
		if err != nil {
			ch <- false
			return
		}

		ch <- true
	}(done, student)

	if channels.OK(done) {
		return &student, nil
	}
	return nil, err
}

func (r *studentRepositoryImpl) FindByUserID(userID uint32) (*models.Student, error) {
	var err error
	var row *sql.Row
	student := models.Student{}
	done := make(chan bool)
	go func(ch chan<- bool, userID uint32) {
		defer close(ch)

		row = r.db.QueryRow(queryutil.GET_STUDENT_BY_USER_ID, userID)
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done, userID)
	

	if channels.OK(done) {

		err = row.Scan(&student.ID, &student.UserID)
		if err != nil {
			return nil, err
		}
		
		return &student, nil
	}

	return nil, err

}
