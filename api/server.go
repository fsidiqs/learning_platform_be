package api

import (
	"fmt"
	"go_jwt_auth/config"

	"go_jwt_auth/api/controllers/coursecontroller"
	"go_jwt_auth/api/controllers/usercontroller"
	"go_jwt_auth/api/controllers/usercontroller/authcontroller"
	"go_jwt_auth/api/middlewares"

	"go_jwt_auth/api/database"
	"go_jwt_auth/api/filestorage/backblaze"

	"go_jwt_auth/api/repository/courserepository"
	"go_jwt_auth/api/repository/studentrepository"
	"go_jwt_auth/api/repository/userrepository"

	"go_jwt_auth/api/router"
	"go_jwt_auth/api/router/routes"

	"log"
	"net/http"

	"github.com/gorilla/handlers"
)

func Run() {
	conf := config.NewAppConfig()
	db, err := database.InitDB(conf.DBConf)
	if err != nil {
		fmt.Println(err)
	}
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "content-type"})
	originsOk := handlers.AllowedOrigins([]string{"http://localhost:3000"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	// ch := handlers.CORS(originsOk, headersOk, methodsOk)
	muxRouter := router.New()

	// make repository
	userRepo := userrepository.NewUserRepository(db)
	studentRepo := studentrepository.NewStudentRepository(db)
	courseRepo := courserepository.NewCourseRepository(db)
	// make controller
	userController := usercontroller.NewUserController(userRepo)
	// initiate routes

	authController := authcontroller.NewAuthController(conf.JWTConf, userRepo, studentRepo)

	// make controller
	fileStorage, err := backblaze.NewBackblaze("courses/", conf.StorageConf)
	if err != nil {

		fmt.Println("error creating filestorage")
		fmt.Printf("%#v\n", err)
	}

	courseController := coursecontroller.NewCourseController(courseRepo, studentRepo, fileStorage)
	// initiate routes

	userRoutes := routes.NewUserRoutes(userController)
	authRoutes := routes.NewAuthRoutes(authController)
	courseRoutes := routes.NewCourseRoutes(courseController)
	// studentRotues := routes.NewStudentCourseRoutes(courseController)
	routeables := []routes.Routables{&userRoutes, &authRoutes, &courseRoutes}

	router := routes.SetupRoutes(muxRouter, routeables, middlewares.AuthConfig{JwtSecret: conf.JWTConf.JWTSecret, UserRepo: userRepo})
	fmt.Printf("\n\tListening on port [::]:%d", conf.APIConf.Port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", conf.APIConf.Port), handlers.CORS(originsOk, headersOk, methodsOk)(router)))

}
