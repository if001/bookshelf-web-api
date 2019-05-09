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
	//FindDescription(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	//CreateDescription(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
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
	defer r.Body.Close()
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
	err = json.NewEncoder(w).Encode(Response{resultCode:200, Content:book})
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

//func (b *bookHandler) FindDescription(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
//	bookId,err := strconv.ParseInt(ps.ByName("book"),10,64)
//	if err != nil {
//		ErrorHandler(service.InternalServerError(err), w ,r)
//		return
//	}
//	descriptions, err := b.BookUseCase.DescriptionFindUseCase(bookId)
//	if err != nil {
//		ErrorHandler(err, w ,r)
//		return
//	}
//	err = json.NewEncoder(w).Encode(Response{resultCode:200, Content:descriptions})
//	if err != nil {
//		ErrorHandler(err, w ,r)
//		return
//	}
//}
//
//
//func (b *bookHandler) CreateDescription(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
//	descriptionRequest := model.DescriptionRequest{}
//	bookId,err := strconv.ParseInt(ps.ByName("book"),10,64)
//	if err != nil {
//		ErrorHandler(service.InternalServerError(err), w ,r)
//		return
//	}
//	body, err := ioutil.ReadAll(r.Body)
//	if err != nil {
//		ErrorHandler(service.BadRequest(err), w, r)
//		return
//	}
//	err = json.Unmarshal(body, &descriptionRequest)
//	if err != nil {
//		ErrorHandler(err, w, r)
//		return
//	}
//	descriptions, err := b.BookUseCase.DescriptionCreateUseCase(bookId, descriptionRequest.Description)
//	if err != nil {
//		ErrorHandler(err, w, r)
//		return
//	}
//	err = json.NewEncoder(w).Encode(Response{resultCode: 200, Content: descriptions})
//	if err != nil {
//		ErrorHandler(err, w, r)
//		return
//	}
//}