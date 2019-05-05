package handler

type ApiHandler interface {
	CategoryHandler
	BookHandler
	AccountHandler
}

type apiHandler struct {
	CategoryHandler
	BookHandler
	AccountHandler
}

func NewApiHandler(uh CategoryHandler, bh BookHandler,ah AccountHandler) ApiHandler {
	return &apiHandler{uh, bh, ah}
}