package handler

import (
	"bookshelf-web-api/application/usecase"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"fmt"
)

type BookHandler interface {
	GetBooks(w http.ResponseWriter, r *http.Request, _ httprouter.Params)
	//FindBook(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	//CreateBook(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	//UpdateBook(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	//FindDescription(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	//CreateDescription(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
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
	ctx := r.Context()
	fmt.Println("getbooks handler")
	books, err := b.BookUseCase.BookListUseCase(ctx)
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
//
//func (b *bookHandler) FindBook(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
//	bookId,err := strconv.ParseInt(ps.ByName("book"),10,64)
//	account := r.Context().Value("account").(*model.Account)
//	if err != nil {
//		ErrorHandler(service.InternalServerError(err), w ,r)
//		return
//	}
//	book, err := b.BookUseCase.BookFindUseCase(bookId, *account)
//	if err != nil {
//		ErrorHandler(err, w ,r)
//		return
//	}
//	err = json.NewEncoder(w).Encode(Response{resultCode:200, Content:book})
//	if err != nil {
//		ErrorHandler(err, w ,r)
//		return
//	}
//}
//
//
//func (b *bookHandler) CreateBook(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
//	bookRequest := model.BookRequest{}
//	defer r.Body.Close()
//	account := r.Context().Value("account").(*model.Account)
//
//	body, err := ioutil.ReadAll(r.Body)
//	if err != nil {
//		ErrorHandler(service.BadRequest(err), w ,r)
//		return
//	}
//	err = json.Unmarshal(body, &bookRequest)
//	if err != nil {
//		ErrorHandler(service.BadRequest(err), w ,r)
//		return
//	}
//	newBook, err := b.BookUseCase.CreateBook(bookRequest, *account)
//	err = json.NewEncoder(w).Encode(Response{resultCode:200, Content:newBook})
//	if err != nil {
//		ErrorHandler(err, w ,r)
//		return
//	}
//}
//
//func (b *bookHandler) UpdateBook(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
//	bookRequest := model.BookRequest{}
//	bookId,err := strconv.ParseInt(ps.ByName("book"),10,64)
//
//	defer r.Body.Close()
//	account := r.Context().Value("account").(*model.Account)
//
//	body, err := ioutil.ReadAll(r.Body)
//	if err != nil {
//		ErrorHandler(service.BadRequest(err), w ,r)
//		return
//	}
//	err = json.Unmarshal(body, &bookRequest)
//	if err != nil {
//		ErrorHandler(service.BadRequest(err), w ,r)
//		return
//	}
//	newBook, err := b.BookUseCase.UpdateBook(bookId, bookRequest, *account)
//	err = json.NewEncoder(w).Encode(Response{resultCode:200, Content:newBook})
//	if err != nil {
//		ErrorHandler(err, w ,r)
//		return
//	}
//}
//
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