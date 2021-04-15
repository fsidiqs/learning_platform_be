package userrepository

import (
	"database/sql"
	"fmt"
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
	var res sql.Result
	done := make(chan bool)
	go func(ch chan<- bool, user models.User) {
		res, err = r.db.Exec(queryutil.INSERT_NEW_USER, user.Name, user.Email, user.Password, user.RoleID)
		
		if err != nil {
			ch <- false
			return
		}
	
		ch <- true
		}(done, user)
		
	if channels.OK(done) {
		lastID, _ := res.LastInsertId()
		user.ID = uint32(lastID)
		return user, nil
	}
	return models.User{}, err
}

func (r *userRepositoryImpl) FindAll() ([]models.User, error) {
	var err error
	var rows *sql.Rows
	users := []models.User{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)

		rows, err = r.db.Query(queryutil.GET_ALL_USER)
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)

	if channels.OK(done) {
		for rows.Next() {
			user := models.User{}
			err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)

			if err != nil {
				return []models.User{}, err
			}

			users = append(users, user)
		}
		err = rows.Err() // get any error encountered in iteration
		if err != nil {
			return []models.User{}, err
		}

		return users, nil
	}
	return []models.User{}, err

}

func (r *userRepositoryImpl) FindById(uid uint32) (models.User, error) {
	var err error
	var row *sql.Row
	user := models.User{}
	done := make(chan bool)
	go func(ch chan<- bool, uid uint32) {
		defer close(ch)

		row = r.db.QueryRow(queryutil.GET_USER_BY_ID, uid)
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done, uid)

	if channels.OK(done) {
		err = row.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
		if err != nil {
			return models.User{}, err
		}
		return user, nil
	}

	return models.User{}, err

}

type queryArgs func() (string, interface{})

func (r *userRepositoryImpl) Update(uid uint32, newValUser models.UserUpdate) (int64, error) {
	var err error
	var res sql.Result

	done := make(chan bool)
	go func(ch chan<- bool, uid uint32, newValUser models.UserUpdate) {
		defer close(ch)
		res, err = r.db.Exec(queryutil.UPDATE_USER_BY_ID, newValUser.Name, newValUser.Email, uid)
		if err != nil {
			fmt.Println(err)
			ch <- false
			return
		}
		ch <- true
	}(done, uid, newValUser)

	if channels.OK(done) {
		return res.RowsAffected()
	}
	return 0, err
}

func (r *userRepositoryImpl) Deactivate(uid uint32) (int64, error) {
	var err error
	var res sql.Result

	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		res, err = r.db.Exec(queryutil.DEACTIVATE_USER, uid)
		if err != nil {
			fmt.Println(err)
			ch <- false
			return
		}
		ch <- true
	}(done)

	if channels.OK(done) {
		return res.RowsAffected()
	}
	return 0, err
}

func (r *userRepositoryImpl) Activate(uid uint32) (int64, error) {
	var err error
	var res sql.Result

	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		res, err = r.db.Exec(queryutil.ACTIVATE_USER, uid)
		if err != nil {
			fmt.Println(err)
			ch <- false
			return
		}
		ch <- true
	}(done)

	if channels.OK(done) {
		return res.RowsAffected()
	}
	return 0, err
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

func (r *userRepositoryImpl) FindByEmail(email string) (*models.User, error) {
	user := &models.User{}
	var err error
	var row *sql.Row

	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)

		row = r.db.QueryRow(queryutil.GET_USER_BY_EMAIL, email)
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)

	if channels.OK(done) {
		err = row.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
		if err != nil {
			return nil, err
		}
		return user, nil
	}
	return nil, err

}

func (r *userRepositoryImpl) FindByEmailWithPassword(email string) (*models.User, error) {
	user := &models.User{}
	var err error
	var row *sql.Row

	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)

		row = r.db.QueryRow(queryutil.GET_USER_BY_EMAIL_WITH_PASSWORD, email)
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)

	if channels.OK(done) {
		err = row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
		if err != nil {
			return nil, err
		}
		return user, nil
	}
	return nil, err

}
