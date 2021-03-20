package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type DBConf struct {
	DBUser               string
	DBPass               string
	DBHost               string
	DBPort               string
	DBName               string
	AllowNativePasswords bool
}

type APIConf struct {
	Port int
}

type JwtConf struct {
	JWTSecret []byte
}

type StorageConf struct {
	BUCKETID   string
	BUCKETNAME string
	ENDPOINT   string
	REGION     string
	KEYID      string
	APPKEY     string
}

type Conf struct {
	APIConf     APIConf
	DBConf      DBConf
	JWTConf     JwtConf
	StorageConf StorageConf
}

func NewAppConfig() Conf {
	var err error
	godotenv.Load()

	if err != nil {
		log.Fatal(err)
	}

	PORT, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Println(err)
		PORT = 9000
	}

	// DBURL = fmt.Sprintf("%s:%s@/%s?charset=utf8mb4&parseTime=True&loc=Local",
	// 	os.Getenv("DB_USER"),
	// 	os.Getenv("DB_PASS"),
	// 	os.Getenv("DB_NAME"))

	DB_USER := os.Getenv("DB_USER")
	DB_PASS := os.Getenv("DB_PASS")
	DB_NAME := os.Getenv("DB_NAME")
	DB_HOST := os.Getenv("DB_HOST")
	DB_PORT := os.Getenv("DB_PORT")

	JWTSECRET := []byte(os.Getenv("JWTSECRET"))

	B2BUCKETID := os.Getenv("B2BUCKETID")
	B2BUCKETNAME := os.Getenv("B2BUCKETNAME")
	B2ENDPOINT := os.Getenv("B2ENDPOINT")
	B2REGION := os.Getenv("B2REGION")
	B2KEYID := os.Getenv("B2KEYID")
	B2APPKEY := os.Getenv("B2APPKEY")
	return Conf{
		APIConf: APIConf{
			Port: PORT,
		},
		DBConf: DBConf{
			DBUser:               DB_USER,
			DBPass:               DB_PASS,
			DBName:               DB_NAME,
			DBHost:               DB_HOST,
			DBPort:               DB_PORT,
			AllowNativePasswords: true,
		},
		JWTConf: JwtConf{
			JWTSecret: JWTSECRET,
		},
		StorageConf: StorageConf{
			BUCKETID:   B2BUCKETID,
			BUCKETNAME: B2BUCKETNAME,
			ENDPOINT:   B2ENDPOINT,
			REGION:     B2REGION,
			KEYID:      B2KEYID,
			APPKEY:     B2APPKEY,
		},
	}
}
