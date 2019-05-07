package handler

import (
	"bookshelf-web-api/application/usecase"
	"bookshelf-web-api/domain/model"
	"bookshelf-web-api/domain/service"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
	"strconv"
)

type BookHandler interface {
	GetBooks(w http.ResponseWriter, r *http.Request, _ httprouter.Params)
	FindBook(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	CreateBook(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	UpdateBook(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
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

func (b *bookHandler) GetBooks(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	account := r.Context().Value("account").(*model.Account)
	books, err := b.BookUseCase.BookListUseCase(*account)
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
	account := r.Context().Value("account").(*model.Account)
	if err != nil {
		ErrorHandler(service.InternalServerError(err), w ,r)
	} else {
		book, err := b.BookUseCase.BookFindUseCase(bookId, *account)
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

var bookRequest model.BookRequest
func (b *bookHandler) CreateBook(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	defer r.Body.Close()
	account := r.Context().Value("account").(*model.Account)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		ErrorHandler(service.BadRequest(err), w ,r)
	}else {
		err = json.Unmarshal(body, &bookRequest)
		if err != nil {
			ErrorHandler(service.BadRequest(err), w ,r)
		} else {
			newBook, err := b.BookUseCase.CreateBook(bookRequest, *account)
			err = json.NewEncoder(w).Encode(Response{resultCode:200, Content:newBook})
			if err != nil {
				ErrorHandler(err, w ,r)
			}
		}
	}
}

func (b *bookHandler) UpdateBook(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	bookId,err := strconv.ParseInt(ps.ByName("book"),10,64)

	defer r.Body.Close()
	account := r.Context().Value("account").(*model.Account)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		ErrorHandler(service.BadRequest(err), w ,r)
	}else {
		err = json.Unmarshal(body, &bookRequest)
		if err != nil {
			ErrorHandler(service.BadRequest(err), w ,r)
		} else {
			newBook, err := b.BookUseCase.UpdateBook(bookId, bookRequest, *account)
			err = json.NewEncoder(w).Encode(Response{resultCode:200, Content:newBook})
			if err != nil {
				ErrorHandler(err, w ,r)
			}
		}
	}
}