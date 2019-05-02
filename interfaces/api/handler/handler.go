package handler

type ApiHandler interface {
	CategoryHandler
	BookHandler
}

type apiHandler struct {
	CategoryHandler
	BookHandler
}

func NewApiHandler(uh CategoryHandler, bh BookHandler) ApiHandler {
	return &apiHandler{uh, bh}
}