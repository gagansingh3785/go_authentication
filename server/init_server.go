package server

import (
	"github.com/gagansingh3785/go_authentication/handlers"
	"github.com/gagansingh3785/go_authentication/middleware"
	"github.com/gorilla/mux"
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
	router.Methods("Get").Path("/home").HandlerFunc(middleware.IsLoggedIn(handlers.Home))
	router.Methods("Options").HandlerFunc(handlers.CorsHandler)
	router.Methods("Post").Path("/generate_session").HandlerFunc(handlers.GenerateSessionHandler)
	router.Methods("Post").Path("/logout").HandlerFunc(middleware.AuthoriseSession(handlers.Logout))
	router.Methods("Post").Path("/write").HandlerFunc(middleware.AuthoriseSession(handlers.Write))
	router.Methods("Get").Path("/detail/{articleID}").HandlerFunc(middleware.IsLoggedIn(handlers.GetDetail))
	router.Methods("Get").Path("/detail/{articleID}/comments").HandlerFunc(handlers.GetArticleComments)
	router.Methods("Post").Path("/{articleID}/comment").HandlerFunc(middleware.AuthoriseSession(handlers.PostArticleComment))
	router.Methods("Post").Path("/{articleID}/like").HandlerFunc(middleware.AuthoriseSession(handlers.LikeArticle))
	router.Methods("Post").Path("/{articleID}/unlike").HandlerFunc(middleware.AuthoriseSession(handlers.UnlikeArticle))
	router.Methods("Get").Path("/{articleID}/is_liked").HandlerFunc(middleware.AuthoriseSession(handlers.IsLikedArticle))
}
