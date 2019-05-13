package handler

import (
	"bookshelf-web-api/domain/service"
	"log"
	"net/http"
)

func BadRequestHandler(err *service.BadRequest, w http.ResponseWriter, r *http.Request) {
	http.Error(w, err.Error(), err.Code)
	log.Println(err, r)
}
func InternalServerErrorHandler(err *service.InternalServerError, w http.ResponseWriter, r *http.Request) {
	http.Error(w, err.Error(), err.Code)
	log.Println(err, r)
}
func RecodeNotFoundErrorHandler(err *service.RecodeNotFoundError, w http.ResponseWriter, r *http.Request) {
	http.Error(w, err.Error(), err.Code)
	log.Println(err, r)
}


func ErrorHandler(err error,w http.ResponseWriter, r *http.Request) {
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