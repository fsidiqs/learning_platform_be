package courserepository

import (
	"database/sql"
	"go_jwt_auth/api/models"
	"go_jwt_auth/api/utils/channels"
)

type courseRepositoryImpl struct {
	db *sql.DB
}

func NewCourseRepository(db *sql.DB) *courseRepositoryImpl {
	return &courseRepositoryImpl{db}
}

func (r *courseRepositoryImpl) FindAll() ([]models.Course, error) {
	var err error
	courses := []models.Course{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)

		// err = r.db.Debug().Model(&models.Course{}).Limit(100).Find(&courses).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)

	if channels.OK(done) {
		return courses, nil
	}
	return []models.Course{}, err

}

func (r *courseRepositoryImpl) Save(course models.Course) (models.Course, error) {
	var err error
	done := make(chan bool)
	go func(ch chan<- bool) {
		// err = r.db.Model(&models.Course{}).Create(&course).Error
		if err != nil {
			ch <- false
			return
		}
		// err = tx.Model(&models.VideoLecture{}).Create()
		ch <- true
	}(done)

	if channels.OK(done) {
		return course, nil
	}
	return models.Course{}, err
}

// func (r *courseRepositoryImpl) TxSave(tx *gorm.DB, course models.Course) (*gorm.DB, models.Course, error) {
// 	var err error

// 	done := make(chan bool)
// 	go func(ch chan<- bool) {
// 		tx := r.db.Begin()
// 		err = tx.Model(&models.Course{}).Create(&course).Error
// 		if err != nil {
// 			tx.Rollback()
// 			ch <- false
// 			return
// 		}
// 		tx.Commit()
// 		ch <- true
// 	}(done)

// 	if channels.OK(done) {
// 		return tx, course, nil
// 	}
// 	return tx, models.Course{}, err
// }

// func (r *courseRepositoryImpl) SaveVideoLecture(lecture models.VideoLecture) (models.VideoLecture, error) {
// 	var err error
// 	done := make(chan bool)
// 	go func(ch chan<- bool) {
// 		err = r.db.Model(&models.VideoLecture{}).Create(&lecture).Error
// 		if err != nil {
// 			ch <- false
// 			return
// 		}
// 		ch <- true
// 	}(done)

// 	if channels.OK(done) {
// 		return lecture, nil
// 	}
// 	return models.VideoLecture{}, err
// }

// func (r *courseRepositoryImpl) TxSaveVideoLecture(tx *gorm.DB, lecture models.VideoLecture) (*gorm.DB, models.VideoLecture, error) {
// 	var err error
// 	done := make(chan bool)
// 	go func(ch chan<- bool) {
// 		tx := r.db.Begin()
// 		err = tx.Model(&models.VideoLecture{}).Create(&lecture).Error
// 		if err != nil {
// 			tx.Rollback()
// 			ch <- false
// 			return
// 		}
// 		tx.Commit()
// 		// err = tx.Model(&models.VideoLecture{}).Create()
// 		ch <- true
// 	}(done)

// 	if channels.OK(done) {
// 		return tx, lecture, nil
// 	}
// 	return tx, models.VideoLecture{}, err
// }
