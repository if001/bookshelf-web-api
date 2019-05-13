package handler

import (
	"bookshelf-web-api/application/usecase"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type BookHandler interface {
	GetBooks(w http.ResponseWriter, r *http.Request, _ httprouter.Params)
	FindBook(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	CreateBook(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	UpdateBook(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	GetBookStatus(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	StartReadBook(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	EndReadBook(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
}

type bookHandler struct {
	BookUseCase usecase.BookUseCase
	AccountUseCase usecase.AccountUseCase
}

func NewBookHandler(b usecase.BookUseCase, a usecase.AccountUseCase) BookHandler {
	return &bookHandler{
		BookUseCase: b,
		AccountUseCase: a,
	}
}

func (b *bookHandler) GetBooks(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	account, err := b.AccountUseCase.GetAccountUseCase(r.Context())
	if err != nil {
		ErrorHandler(err, w ,r)
		return
	}
	books, err := b.BookUseCase.BookListUseCase(*account)
	if err != nil {
		ErrorHandler(err, w ,r)
		return
	}
	err = json.NewEncoder(w).Encode(Response{resultCode:200, Content:books})
	if err != nil {
		ErrorHandler(err, w ,r)
		return
	}
}

func (b *bookHandler) CreateBook(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	defer func() {
		err := r.Body.Close()
		if err != nil {
			ErrorHandler(err, w ,r)
			return
		}
	} ()

	account, err := b.AccountUseCase.GetAccountUseCase(r.Context())
	if err != nil {
		ErrorHandler(err, w ,r)
		return
	}
	body := r.Body
	defer r.Body.Close()
	book, err := b.BookUseCase.BookRequestBind(body)
	if err != nil {
		ErrorHandler(err, w ,r)
		return
	}

	newBook, err := b.BookUseCase.CreateBook(*book, *account)
	if err != nil {
		ErrorHandler(err, w ,r)
		return
	}
	err = json.NewEncoder(w).Encode(Response{resultCode:200, Content:newBook})
	if err != nil {
		ErrorHandler(err, w ,r)
		return
	}
}




func (b *bookHandler) FindBook(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	bookId,err := strconv.ParseInt(ps.ByName("book"),10,64)
	if err != nil {
		ErrorHandler(err, w ,r)
		return
	}
	account, err := b.AccountUseCase.GetAccountUseCase(r.Context())
	if err != nil {
		ErrorHandler(err, w ,r)
		return
	}
	book, err := b.BookUseCase.BookFindUseCase(bookId, *account)
	if err != nil {
		ErrorHandler(err, w ,r)
		return
	}
	bookResponse := b.BookUseCase.ModelToResponse(*book)
	err = json.NewEncoder(w).Encode(Response{resultCode:200, Content:bookResponse})
	if err != nil {
		ErrorHandler(err, w ,r)
		return
	}
}

func (b *bookHandler) UpdateBook(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	bookId,err := strconv.ParseInt(ps.ByName("book"),10,64)
	if err != nil {
		ErrorHandler(err, w ,r)
		return
	}

	account, err := b.AccountUseCase.GetAccountUseCase(r.Context())
	if err != nil {
		ErrorHandler(err, w ,r)
		return
	}

	body := r.Body
	defer r.Body.Close()
	book, err := b.BookUseCase.BookRequestBind(body)
	if err != nil {
		ErrorHandler(err, w ,r)
		return
	}
	book.ID = bookId
	newBook, err := b.BookUseCase.UpdateBook(*book, *account)
	if err != nil {
		ErrorHandler(err, w ,r)
		return
	}
	err = json.NewEncoder(w).Encode(Response{resultCode:200, Content:newBook})
	if err != nil {
		ErrorHandler(err, w ,r)
		return
	}
}

func (b *bookHandler) GetBookStatus(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	bookId, err := strconv.ParseInt(ps.ByName("book"),10,64)
	if err != nil {
		ErrorHandler(err, w ,r)
		return
	}
	account, err := b.AccountUseCase.GetAccountUseCase(r.Context())
	if err != nil {
		ErrorHandler(err, w ,r)
		return
	}
	bookStatusResponse, err := b.BookUseCase.GetBookState(bookId, *account)
	if err != nil {
		ErrorHandler(err, w ,r)
		return
	}

	err = json.NewEncoder(w).Encode(Response{resultCode:200, Content:bookStatusResponse})
	if err != nil {
		ErrorHandler(err, w ,r)
		return
	}
}

func (b *bookHandler) StartReadBook(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	bookId, err := strconv.ParseInt(ps.ByName("book"),10,64)
	if err != nil {
		ErrorHandler(err, w ,r)
		return
	}
	account, err := b.AccountUseCase.GetAccountUseCase(r.Context())
	if err != nil {
		ErrorHandler(err, w ,r)
		return
	}

	book, err := b.BookUseCase.StartReadBook(bookId,  *account)
	if err != nil {
		ErrorHandler(err, w ,r)
		return
	}

	err = json.NewEncoder(w).Encode(Response{resultCode:200, Content:book})
	if err != nil {
		ErrorHandler(err, w ,r)
		return
	}
}

func (b *bookHandler) EndReadBook(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	bookId, err := strconv.ParseInt(ps.ByName("book"),10,64)
	if err != nil {
		ErrorHandler(err, w ,r)
		return
	}
	account, err := b.AccountUseCase.GetAccountUseCase(r.Context())
	if err != nil {
		ErrorHandler(err, w ,r)
		return
	}

	book, err := b.BookUseCase.EndReadBook(bookId,  *account)
	if err != nil {
		ErrorHandler(err, w ,r)
		return
	}

	err = json.NewEncoder(w).Encode(Response{resultCode:200, Content:book})
	if err != nil {
		ErrorHandler(err, w ,r)
		return
	}
}