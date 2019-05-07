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

	router.GET("/books", middlewares.Then(h.BookList))
	router.POST("/books", middlewares.Then(h.Create))

	router.GET("/book/:book", middlewares.Then(h.FindBook))
	router.PUT("/book/:book", middlewares.Then(h.Update))
	// router.DELETE("/book/:book", middlewares.Then(h.BookList))

	router.GET("/book/:book/description", middlewares.Then(h.FindDescription))
	// router.POST("/book/:book/description", middlewares.Then(h.FindDescription))
	// router.PUT("/description/:description", middlewares.Then(h.FindDescription))

	return  router
}
