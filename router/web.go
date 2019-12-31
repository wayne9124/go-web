package router

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"ikdev/smartcherry/config"
	"ikdev/smartcherry/controller"
	"ikdev/smartcherry/middleware"
	"net/http"
)

var db *gorm.DB
var conf config.Conf

func WebRouter(conn *gorm.DB, cgf config.Conf) *mux.Router {
	db = conn
	conf = cgf

	router := mux.NewRouter()
	router.HandleFunc("/", homeAction).Methods("GET")
	router.HandleFunc("/users", newUserAction).Methods("POST")
	router.HandleFunc("/login", loginAction).Methods("POST")

	router.Use(middleware.LoggingMiddleware)

	authRouter := router.PathPrefix("/admin").Subrouter()
	authRouter.HandleFunc("/test", testUserAction)
	authRouter.Use(middleware.AuthMiddleware)

	return router
}

func homeAction(w http.ResponseWriter, r *http.Request) {
	homeController := controller.HomeController{
		Controller: controller.Controller{
			DB:       db,
			Response: w,
			Request:  r,
			Config:   conf,
		}}

	homeController.Show()
}

func newUserAction(w http.ResponseWriter, r *http.Request) {
	userController := controller.UserController{Controller: controller.Controller{
		DB:       db,
		Response: w,
		Request:  r,
		Config:   conf,
	}}

	userController.Insert()
}

func testUserAction(w http.ResponseWriter, r *http.Request) {
	userController := controller.UserController{Controller: controller.Controller{
		DB:       db,
		Response: w,
		Request:  r,
		Config:   conf,
	}}

	userController.Profile()
}

func loginAction(w http.ResponseWriter, r *http.Request) {
	authController := controller.AuthController{Controller: controller.Controller{
		DB:       db,
		Response: w,
		Request:  r,
		Config:   conf,
	}}

	authController.Login()
}
