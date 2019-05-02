package handler

import (
	"bookshelf-web-api/application/usecase"
	"encoding/json"
	"net/http"
)

type BookHandler interface {
	BookList(w http.ResponseWriter, r *http.Request)
}

type bookHandler struct {
	BookUseCase usecase.BookUseCase
}

func NewBookHandler(b usecase.BookUseCase) BookHandler {
	return &bookHandler{
		BookUseCase: b,
	}
}

func (b *bookHandler) BookList(w http.ResponseWriter, r *http.Request) {
	books, err := b.BookUseCase.BookListUseCase()
	if err != nil {
		ErrorHandler(err, w ,r)
	}
	err = json.NewEncoder(w).Encode(Response{resultCode:200, Content:books})
	if err != nil {
		ErrorHandler(err, w ,r)
	}
}
	err = json.NewEncoder(w).Encode(Response{resultCode:200, Content:book})
	if err != nil {
		ErrorHandler(err, w ,r)
	}
}