package handler

import (
	"bookshelf-web-api/application/usecase"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"strconv"
)

type DescriptionHandler interface {
	FindDescription(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	// UpdateDescription(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
}

type descriptionHandler struct {
	DescriptionUseCase usecase.DescriptionUseCase
}

func NewDescriptionHandler(u usecase.DescriptionUseCase) DescriptionHandler {
	return &descriptionHandler{
		DescriptionUseCase: u,
	}
}

func bindForm(body io.ReadCloser, form interface{}) error {
	// todo バインドの処理が怪しい
	// todo bodyにformのkeyにない場合は空のstructが変えるのcatchできてない
	readBody, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(readBody, form)
	if err != nil {
		return err
	}
	return nil
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


//func (d *descriptionHandler) UpdateDescription(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
//	descriptionRequest := model.DescriptionRequest{}
//
//	descriptionId,err := strconv.ParseInt(ps.ByName("description"),10,64)
//	if err != nil {
//		ErrorHandler(service.InternalServerError(err), w ,r)
//		return
//	}
//	err = bindForm(r.Body, &descriptionRequest)
//	if err != nil {
//		ErrorHandler(service.BadRequest(err), w ,r)
//		return
//	}
//
//	if descriptionRequest.Description == "" {
//		ErrorHandler(service.BadRequest(errors.New("bind error")), w ,r)
//		return
//	}
//	newDescription, err := d.DescriptionUseCase.DescriptionUpdateUseCase(descriptionId, descriptionRequest)
//	err = json.NewEncoder(w).Encode(Response{resultCode: 200, Content: newDescription})
//	if err != nil {
//		ErrorHandler(err, w, r)
//		return
//	}
//}