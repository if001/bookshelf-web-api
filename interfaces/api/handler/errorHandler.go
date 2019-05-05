package handler

import (
	"bookshelf-web-api/domain/service"
	"log"
	"net/http"
)

func BadRequestHandler(err interface{}, w http.ResponseWriter, r *http.Request) {
	http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	log.Println(http.StatusText(http.StatusBadGateway), r)
}
func InternalServerErrorHandler(err interface{}, w http.ResponseWriter, r *http.Request) {
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	log.Println(http.StatusText(http.StatusInternalServerError), r)
}

func RecodeNotFoundErrorHandler(err interface{}, w http.ResponseWriter, r *http.Request) {
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	log.Println(http.StatusText(http.StatusNotFound), r)
}

func ErrorHandler(err interface{},w http.ResponseWriter, r *http.Request) {
	// TODO 型で分岐じゃなくてinterfaceで分岐に変える
	switch err.(type) {
	case service.BadRequest:
		BadRequestHandler(err, w, r)
	case service.InternalServerError:
		InternalServerErrorHandler(err, w, r)
	case service.RecodeNotFoundError:
		RecodeNotFoundErrorHandler(err, w, r)
	default:
		InternalServerErrorHandler(err, w, r)
	}
}