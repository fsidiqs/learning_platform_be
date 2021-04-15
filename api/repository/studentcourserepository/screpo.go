package studentcourserepository

import (
	"database/sql"
	
)

type studentCourseRepositoryImpl struct {
	db *sql.DB
	// tx *sql.Tx
}

func NewCourseRepository(db *sql.DB) *studentCourseRepositoryImpl {
	return &studentCourseRepositoryImpl{db}
}

