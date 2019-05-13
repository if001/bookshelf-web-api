package handler

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"bookshelf-web-api/application/usecase"
	"encoding/json"
	"strconv"
)

type Response struct {
	resultCode uint
	Content interface{}
}


type CategoryHandler interface {
	CategoryList(w http.ResponseWriter, r *http.Request,_ httprouter.Params)
	CategoryLogicalDelete(w http.ResponseWriter, r *http.Request,_ httprouter.Params)
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
	category, err := u.CategoryUseCase.CategoryUseCase()
	if err != nil {
		ErrorHandler(err, w, r)
		return
	}
	err = json.NewEncoder(w).Encode(Response{resultCode:200, Content:category})
	if err != nil {
		ErrorHandler(err, w, r)
		return
	}
}

func (u *categoryHandler) CategoryLogicalDelete(w http.ResponseWriter, r *http.Request,ps httprouter.Params) {
	bookId,err := strconv.ParseInt(ps.ByName("book"),10,64)
	if err != nil {
		ErrorHandler(err, w, r)
		return
	}
	categoryId,err := strconv.ParseInt(ps.ByName("category"),10,64)
	if err != nil {
		ErrorHandler(err, w, r)
		return
	}
	err = u.CategoryUseCase.CategoryLogicalDeleteCase(bookId, categoryId)
	if err != nil {
		ErrorHandler(err, w, r)
		return
	}
}