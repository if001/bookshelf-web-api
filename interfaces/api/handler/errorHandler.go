package handler

import (
	"bookshelf-web-api/domain/service"
	"log"
	"net/http"
)

func BadRequestHandler(err *service.BadRequest, w http.ResponseWriter, r *http.Request) {
	code := http.StatusBadRequest
	http.Error(w, http.StatusText(code), code)
	log.Println(err, r)
}
func InternalServerErrorHandler(err *service.InternalServerError, w http.ResponseWriter, r *http.Request) {
	code := http.StatusInternalServerError
	http.Error(w, http.StatusText(code), code)
	log.Println(err, r)
}
func RecodeNotFoundErrorHandler(err *service.RecodeNotFoundError, w http.ResponseWriter, r *http.Request) {
	code := http.StatusNotFound
	http.Error(w, http.StatusText(code), code)
	log.Println(err, r)
}


func ErrorHandler(err error, w http.ResponseWriter, r *http.Request) {
	switch err := err.(type) {
	case *service.BadRequest:
		BadRequestHandler(err, w, r)
	case *service.InternalServerError:
		InternalServerErrorHandler(err, w, r)
	case *service.RecodeNotFoundError:
		RecodeNotFoundErrorHandler(err, w, r)
	default:
		http.Error(w, "Not Found", 400)
	}
}