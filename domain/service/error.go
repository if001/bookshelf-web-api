package service

type BadRequest struct {
	Message string
}
func (err *BadRequest) Error() string { return err.Message }
func NewBadRequest(err error) *BadRequest {
	return &BadRequest{err.Error()}
}


type InternalServerError struct {
	Message string
}
func (err *InternalServerError) Error() string { return err.Message }
func NewInternalServerError(err error) *InternalServerError {
	return &InternalServerError{err.Error()}
}


type RecodeNotFoundError struct {
	Message string
}
func (err *RecodeNotFoundError) Error() string { return err.Message }
func NewRecodeNotFoundError(err error) *RecodeNotFoundError {
	return &RecodeNotFoundError{err.Error()}
}
