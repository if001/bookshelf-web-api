package handler

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"bookshelf-web-api/application/usecase"
	"encoding/json"
)

type Response struct {
	resultCode uint
	Content interface{}
}


type CategoryHandler interface {
	CategoryList(w http.ResponseWriter, r *http.Request,_ httprouter.Params)
}

type categoryHandler struct {
	CategoryUseCase usecase.CategoryUseCase
}

func NewCategoryHandler(c usecase.CategoryUseCase) CategoryHandler {
	return &categoryHandler{
		CategoryUseCase: c,
	}
}

func (u *categoryHandler) CategoryList(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	cate, err := u.CategoryUseCase.CategoryUseCase()
	if err != nil {
		ErrorHandler(err, w, r)
		return
	}
	err = json.NewEncoder(w).Encode(Response{resultCode:200, Content:cate})
	if err != nil {
		ErrorHandler(err, w, r)
		return
	}
}