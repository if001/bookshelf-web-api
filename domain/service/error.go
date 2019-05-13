package service

import "net/http"

type BadRequest struct {
	Code int
	Message string
}
func (err *BadRequest) Error() string { return err.Message }
func NewBadRequest() *BadRequest {
	code := http.StatusBadRequest
	return &BadRequest{code, http.StatusText(code)}
}


type InternalServerError struct {
	Code int
	Message string
}
func (err *InternalServerError) Error() string { return err.Message }
func NewInternalServerError() *InternalServerError {
	code := http.StatusInternalServerError
	return &InternalServerError{code, http.StatusText(code)}
}


type RecodeNotFoundError struct {
	Code int
	Message string
}
func (err *RecodeNotFoundError) Error() string { return err.Message }
func NewRecodeNotFoundError() *RecodeNotFoundError {
	code := http.StatusNotFound
	return &RecodeNotFoundError{code, "Record Not Found"}
}
