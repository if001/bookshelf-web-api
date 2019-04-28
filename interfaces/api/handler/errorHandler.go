package handler

import (
	"net/http"
	"fmt"
	"log"
	"bookshelf-web-api/domain/service"
)

func BadRequestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, http.StatusText(http.StatusBadRequest))
	log.Println(http.StatusText(http.StatusBadGateway), r)
}
func InternalServerErrorHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, http.StatusText(http.StatusInternalServerError))
	log.Println(http.StatusText(http.StatusInternalServerError), r)
}

func RecodeNotFoundErrorHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, http.StatusText(http.StatusNotFound))
	log.Println(http.StatusText(http.StatusNotFound), r)
}

func ErrorHandler(err interface{},w http.ResponseWriter, r *http.Request) {
	// TODO 型で分岐じゃなくてinterfaceで分岐に変える
	switch err.(type) {
	case service.BadRequest:
		BadRequestHandler(w, r)
	case service.InternalServerError:
		InternalServerErrorHandler(w, r)
	case service.RecodeNotFoundError:
		RecodeNotFoundErrorHandler(w, r)
	default:
		InternalServerErrorHandler(w, r)
	}
}