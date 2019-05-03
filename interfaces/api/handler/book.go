package handler

import (
	"bookshelf-web-api/application/usecase"
	"bookshelf-web-api/domain/service"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type BookHandler interface {
	BookList(w http.ResponseWriter, r *http.Request, _ httprouter.Params)
	FindBook(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	FindDescription(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
}

type bookHandler struct {
	BookUseCase usecase.BookUseCase
}

func NewBookHandler(b usecase.BookUseCase) BookHandler {
	return &bookHandler{
		BookUseCase: b,
	}
}

func (b *bookHandler) BookList(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	books, err := b.BookUseCase.BookListUseCase()
	if err != nil {
		ErrorHandler(err, w ,r)
	} else {
		err = json.NewEncoder(w).Encode(Response{resultCode:200, Content:books})
		if err != nil {
			ErrorHandler(err, w ,r)
		}
	}
}

func (b *bookHandler) FindBook(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	bookId,err := strconv.ParseInt(ps.ByName("book"),10,64)
	if err != nil {
		ErrorHandler(service.InternalServerError(err), w ,r)
	} else {
		book, err := b.BookUseCase.BookFindUseCase(bookId)
		if err != nil {
			ErrorHandler(err, w ,r)
		} else {
			err = json.NewEncoder(w).Encode(Response{resultCode:200, Content:book})
			if err != nil {
				ErrorHandler(err, w ,r)
			}
		}
	}
}

func (b *bookHandler) FindDescription(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	bookId,err := strconv.ParseInt(ps.ByName("book"),10,64)
	if err != nil {
		ErrorHandler(service.InternalServerError(err), w ,r)
	} else {
		descriptions, err := b.BookUseCase.DescriptionUseCase(bookId)
		if err != nil {
			ErrorHandler(err, w ,r)
		} else {
			err = json.NewEncoder(w).Encode(Response{resultCode:200, Content:descriptions})
			if err != nil {
				ErrorHandler(err, w ,r)
			}
		}
	}
}

