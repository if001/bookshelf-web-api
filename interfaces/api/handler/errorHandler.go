package handler

import (
	"net/http"
	"fmt"
	"log"
	"bookshelf-web-api/domain/service"
)

func BadRequestHandler(err interface{}, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, http.StatusText(http.StatusBadRequest))
	log.Println(http.StatusText(http.StatusBadGateway), r)
	log.Println(err)
}
func InternalServerErrorHandler(err interface{}, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, http.StatusText(http.StatusInternalServerError))
	log.Println(http.StatusText(http.StatusInternalServerError), r)
	log.Println(err)
}

func RecodeNotFoundErrorHandler(err interface{}, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, http.StatusText(http.StatusNotFound))
	log.Println(http.StatusText(http.StatusNotFound), r)
	log.Println(err)
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