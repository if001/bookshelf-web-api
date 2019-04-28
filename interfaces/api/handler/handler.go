package handler

type ApiHandler interface {
	CategoryHandler
}

type apiHandler struct {
	CategoryHandler
}

func NewApiHandler(uh CategoryHandler) ApiHandler {
	return &apiHandler{uh}
}