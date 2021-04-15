package database

import (
	"database/sql"
	"fmt"
	"go_jwt_auth/config"
	"log"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

func InitDB(c config.DBConf) (*sql.DB, error) {
	var err error
	cfg := &mysql.Config{
		User:                 c.DBUser,
		Passwd:               c.DBPass,
		Net:                  "tcp",
		Addr:                 fmt.Sprintf("%v:%v", c.DBHost, c.DBPort),
		DBName:               c.DBName,
		AllowNativePasswords: c.AllowNativePasswords,
		ParseTime:            true,
	}
	db, err := sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		log.Panic(err)
	}

	//Ping = check database availability
	if err = db.Ping(); err != nil {
		log.Panic(err)
	}
	log.Println("Connected to the database....")
	return db, nil
}
