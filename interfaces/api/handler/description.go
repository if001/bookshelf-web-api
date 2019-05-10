package handler

import (
	"bookshelf-web-api/application/usecase"
	"encoding/json"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"strconv"
)

type DescriptionHandler interface {
	FindDescription(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	CreateDescription(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	UpdateDescription(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
}

type descriptionHandler struct {
	DescriptionUseCase usecase.DescriptionUseCase
}

func NewDescriptionHandler(u usecase.DescriptionUseCase) DescriptionHandler {
	return &descriptionHandler{
		DescriptionUseCase: u,
	}
}

func (h *descriptionHandler) FindDescription(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	bookId,err := strconv.ParseInt(ps.ByName("book"),10,64)
	if err != nil {
		ErrorHandler(err, w ,r)
		return
	}
	descriptions, err := h.DescriptionUseCase.DescriptionFindUseCase(bookId)
	if err != nil {
		ErrorHandler(err, w ,r)
		return
	}
	err = json.NewEncoder(w).Encode(Response{resultCode:200, Content:descriptions})
	if err != nil {
		ErrorHandler(err, w ,r)
		return
	}
}


func (h *descriptionHandler) CreateDescription(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	bookId, err := strconv.ParseInt(ps.ByName("book"),10,64)
	if err != nil {
		ErrorHandler(err, w ,r)
		return
	}

	body := r.Body
	defer func() {
		err = body.Close()
		if err != nil {
			ErrorHandler(err, w, r)
		}
	}()
	description, err := h.DescriptionUseCase.DescriptionRequestBindWithPath(bookId, body)
	if err != nil {
		ErrorHandler(err, w, r)
		return
	}

	descriptions, err := h.DescriptionUseCase.DescriptionCreateUseCase(*description)
	if err != nil {
		ErrorHandler(err, w, r)
		return
	}
	err = json.NewEncoder(w).Encode(Response{resultCode: 200, Content: descriptions})
	if err != nil {
		ErrorHandler(err, w, r)
		return
	}
}

func (h *descriptionHandler) UpdateDescription(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	descriptionId,err := strconv.ParseInt(ps.ByName("description"),10,64)
	if err != nil {
		ErrorHandler(err, w ,r)
		return
	}

	body := r.Body
	defer func() {
		err = body.Close()
		if err != nil {
			ErrorHandler(err, w, r)
		}
	}()
	descriptionRequest, err :=  h.DescriptionUseCase.DescriptionRequestBind(body)
	if err != nil {
		ErrorHandler(err, w ,r)
		return
	}
	descriptionRequest.ID = descriptionId

	newDescription, err := h.DescriptionUseCase.DescriptionUpdateUseCase(*descriptionRequest)
	if err != nil {
		ErrorHandler(err, w ,r)
		return
	}

	err = json.NewEncoder(w).Encode(Response{resultCode: 200, Content: newDescription})
	if err != nil {
		ErrorHandler(err, w, r)
		return
	}
}