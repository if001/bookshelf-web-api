package router

import (
	"bookshelf-web-api/interfaces/api/handler"
	"bookshelf-web-api/interfaces/api/middleware"
	"github.com/julienschmidt/httprouter"
	"net/http"
)


func Route(h handler.ApiHandler) http.Handler {
	middlewares := middleware.NewMws(h.AuthMiddleware)

	router := httprouter.New()

	//router.HandleFunc("/", middlewares.Then(handler.Index))
	router.GET("/books", middlewares.Then(h.BookList))
	router.GET("/categories", h.CategoryList)
	router.GET("/book/:book", h.FindBook)
	router.GET("/book/:book/description", h.FindDescription)
	return  router
}
