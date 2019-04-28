package handler

import (
	"net/http"
	"bookshelf-web-api/application/usecase"
	"encoding/json"
)

type Response struct {
	resultCode uint
	Message string
}


type CategoryHandler interface {
	Hoge(w http.ResponseWriter, r *http.Request)
	Fuga(w http.ResponseWriter, r *http.Request)
}

type categoryHandler struct {
	CategoryUseCase usecase.CategoryUseCase
}

func NewCategoryHandler(c usecase.CategoryUseCase) CategoryHandler {
	return &categoryHandler{
		CategoryUseCase: c,
	}
}

func (u *categoryHandler) Hoge(w http.ResponseWriter, r *http.Request) {
	cate, err := u.CategoryUseCase.CategoryUseCase()
	if err != nil {
		ErrorHandler(err, w, r)
	}
	err = json.NewEncoder(w).Encode(cate)
	if err != nil {
		ErrorHandler(err, w, r)
	}
}

func (u *categoryHandler) Fuga(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("huga"))
}
