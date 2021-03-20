package api

import (
	"fmt"
	"go_jwt_auth/config"

	"go_jwt_auth/api/controllers/coursecontroller"
	"go_jwt_auth/api/controllers/usercontroller"
	"go_jwt_auth/api/controllers/usercontroller/authcontroller"

	"go_jwt_auth/api/database"
	"go_jwt_auth/api/filestorage/backblaze"

	"go_jwt_auth/api/repository/courserepository"
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

	ch := handlers.CORS(handlers.AllowedOrigins([]string{"*"}))
	muxRouter := router.New()

	// make repository
	userRepo := userrepository.NewUserRepository(db)
	// make controller
	userController := usercontroller.NewUserController(userRepo)
	// initiate routes
	userRoutes := routes.NewUserRoutes(userController)

	authController := authcontroller.NewAuthController(conf.JWTConf, userRepo)
	authRoutes := routes.NewAuthRoutes(authController)

	courseRepo := courserepository.NewCourseRepository(db)
	// make controller
	fileStorage, err := backblaze.NewBackblaze("courses/", conf.StorageConf)
	if err != nil {

		fmt.Println("error creating filestorage")
		fmt.Printf("%#v\n", err)
	}
	courseController := coursecontroller.NewCourseController(courseRepo, fileStorage)
	// initiate routes
	courseRoutes := routes.NewCourseRoutes(courseController)

	routeables := []routes.Routables{&userRoutes, &authRoutes, &courseRoutes}

	router := routes.SetupRoutes(muxRouter, routeables)
	fmt.Printf("\n\tListening on port [::]:%d", conf.APIConf.Port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", conf.APIConf.Port), ch(router)))

}
