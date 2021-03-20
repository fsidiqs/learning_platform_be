package userrepository

import (
	"database/sql"
	"errors"
	"go_jwt_auth/api/models"
	"go_jwt_auth/api/utils/channels"
	"go_jwt_auth/api/utils/queryutil"

	"gorm.io/gorm"
)

type userRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *userRepositoryImpl {
	return &userRepositoryImpl{db}
}

func (r *userRepositoryImpl) Save(user models.User) (models.User, error) {
	var err error
	done := make(chan bool)
	go func(ch chan<- bool) error {
		_, err := r.db.Exec(queryutil.INSERT_USER, user.Name, user.Email, user.Password)
		if err != nil {
			ch <- false
			return err
		}

		ch <- true
		return nil
	}(done)

	if channels.OK(done) {
		return user, nil
	}
	return models.User{}, err
}

func (r *userRepositoryImpl) FindAll() ([]models.User, error) {
	var err error
	users := []models.User{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)

		// err = r.db.Debug().Model(&models.User{}).Limit(100).Find(&users).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)

	if channels.OK(done) {
		return users, nil
	}
	return []models.User{}, err

}

func (r *userRepositoryImpl) FindById(uid uint32) (models.User, error) {
	var err error
	user := models.User{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)

		// err = r.db.Debug().Model(&models.User{}).Where("id = ?", uid).Take(&user).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)

	if channels.OK(done) {
		return user, nil
	}
	if gorm.ErrRecordNotFound == err {
		return models.User{}, errors.New("user not found")
	}
	return models.User{}, err

}

func (r *userRepositoryImpl) Update(uid uint32, newValUser models.UserUpdate) (int64, error) {
	var res *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		// res = r.db.Debug().Model(&models.UserUpdate{}).Where("id = ?", uid).Updates(newValUser)

		ch <- true
	}(done)
	if channels.OK(done) {
		if res.Error != nil {
			return 0, res.Error
		}
		return res.RowsAffected, nil
	}
	return 0, res.Error
}

func (r *userRepositoryImpl) Delete(uid uint32) (int64, error) {
	var res *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		// res = r.db.Debug().Model(&models.User{}).Where("id = ?", uid).Take(&models.User{}).Delete(&models.User{})
	}(done)
	// unfinished
	return 0, res.Error
}

func (r *userRepositoryImpl) FindByEmail(email string) (models.User, error) {
	user := models.User{}
	var err error

	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)

		// err = r.db.Debug().Model(&models.User{}).Where("email = ?", email).Take(&user).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)

	if channels.OK(done) {
		return user, nil
	}
	if gorm.ErrRecordNotFound == err {
		return models.User{}, errors.New("user not found")
	}
	return models.User{}, err

}
