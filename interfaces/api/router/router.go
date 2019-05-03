package router

import (
	"bookshelf-web-api/interfaces/api/handler"
	"github.com/julienschmidt/httprouter"
	"net/http"
)


func Route(h handler.ApiHandler) http.Handler {
	//middlewares := middleware.NewMws(mysql.DBConnection)

	router := httprouter.New()

	//router.HandleFunc("/", middlewares.Then(handler.Index))
	router.GET("/books", h.BookList)
	router.GET("/categories", h.CategoryList)
	router.GET("/book/:book", h.FindBook)
	router.GET("/book/:book/description", h.FindDescription)
	return  router
}
