// Code generated by go-swagger; DO NOT EDIT.

package books

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "bookshelf-web-api/models"
)

// GetBooksOKCode is the HTTP code returned for type GetBooksOK
const GetBooksOKCode int = 200

/*GetBooksOK Success

swagger:response getBooksOK
*/
type GetBooksOK struct {

	/*
	  In: Body
	*/
	Payload *models.BooksResponse `json:"body,omitempty"`
}

// NewGetBooksOK creates GetBooksOK with default headers values
func NewGetBooksOK() *GetBooksOK {

	return &GetBooksOK{}
}

// WithPayload adds the payload to the get books o k response
func (o *GetBooksOK) WithPayload(payload *models.BooksResponse) *GetBooksOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get books o k response
func (o *GetBooksOK) SetPayload(payload *models.BooksResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetBooksOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetBooksDefault generic error response

swagger:response getBooksDefault
*/
type GetBooksDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetBooksDefault creates GetBooksDefault with default headers values
func NewGetBooksDefault(code int) *GetBooksDefault {
	if code <= 0 {
		code = 500
	}

	return &GetBooksDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get books default response
func (o *GetBooksDefault) WithStatusCode(code int) *GetBooksDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get books default response
func (o *GetBooksDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get books default response
func (o *GetBooksDefault) WithPayload(payload *models.Error) *GetBooksDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get books default response
func (o *GetBooksDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetBooksDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
