package handler

type ApiHandler interface {
	CategoryHandler
	BookHandler
	AccountHandler
	DescriptionHandler
}

type apiHandler struct {
	CategoryHandler
	BookHandler
	AccountHandler
	DescriptionHandler
}

func NewApiHandler(uh CategoryHandler, bh BookHandler,ah AccountHandler,dh DescriptionHandler) ApiHandler {
	return &apiHandler{uh, bh, ah, dh}
}