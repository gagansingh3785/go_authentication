package server

import (
	"github.com/gagansingh3785/go_authentication/handlers"
	"github.com/gorilla/mux"
)

const (
	some  = "hello"
	other = 123
)

func CreateNewRouter() *mux.Router {
	router := mux.NewRouter()
	registerRoutes(router)
	return router
}

func registerRoutes(router *mux.Router) {
	//General Routes
	router.Methods("Get").Path("/register").HandlerFunc(handlers.Register)
	router.Methods("Post").Path("/register").HandlerFunc(handlers.Register)
	router.Methods("Get").Path("/login").HandlerFunc(handlers.Login)
	router.Methods("Post").Path("/login").HandlerFunc(handlers.HandleLogin)
	router.Methods("Get").Path("/home").HandlerFunc(handlers.Home)
	router.Methods("Post, Options").Path("/send_message").HandlerFunc(handlers.SendMail)
}
