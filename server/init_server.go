package server

import (
	"github.com/gagansingh3785/go_authentication/handlers"
	"github.com/gagansingh3785/go_authentication/middleware"
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
	router.Methods("Post").Path("/register").HandlerFunc(handlers.Register)
	router.Methods("Post").Path("/login").HandlerFunc(handlers.Login)
	router.Methods("Get").Path("/home").HandlerFunc(middleware.AuthoriseSession(handlers.Home))
	router.Methods("Post").Path("/send_message").HandlerFunc(handlers.SendMail)
	router.Methods("Options").HandlerFunc(handlers.CorsHandler)
	router.Methods("Post").Path("/generate_session").HandlerFunc(handlers.GenerateSessionHandler)
}
