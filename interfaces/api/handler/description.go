package handler

import (
	"bookshelf-web-api/application/usecase"
	"bookshelf-web-api/domain/model"
	"bookshelf-web-api/domain/service"
	"encoding/json"
	"errors"
	"github.com/julienschmidt/httprouter"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
)

type DescriptionHandler interface {
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

func (d *descriptionHandler) UpdateDescription(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	descriptionRequest := new(model.DescriptionRequest)

	descriptionId,err := strconv.ParseInt(ps.ByName("description"),10,64)
	if err != nil {
		ErrorHandler(service.InternalServerError(err), w ,r)
		return
	}
	err = bindForm(r.Body, &descriptionRequest)
	if err != nil {
		ErrorHandler(service.BadRequest(err), w ,r)
		return
	}

	if descriptionRequest.Description == "" {
		ErrorHandler(service.BadRequest(errors.New("bind error")), w ,r)
		return
	}
	newDescription, err := d.DescriptionUseCase.DescriptionUpdateUseCase(descriptionId, descriptionRequest.Description)
	err = json.NewEncoder(w).Encode(Response{resultCode: 200, Content: newDescription})
	if err != nil {
		ErrorHandler(err, w, r)
		return
	}
}