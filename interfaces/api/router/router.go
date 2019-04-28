package router

import (
	"bookshelf-web-api/interfaces/api/handler"
	"net/http"
)


func Route(h handler.ApiHandler) http.Handler {
	//middlewares := middleware.NewMws(mysql.DBConnection)

	router := http.NewServeMux()
	//router.HandleFunc("/", middlewares.Then(handler.Index))
	router.HandleFunc("/huga", h.Fuga)
	router.HandleFunc("/hoge", h.Hoge)

	return  router
}
