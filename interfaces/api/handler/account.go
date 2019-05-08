package handler

import (
	"bookshelf-web-api/application/usecase"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)



type AccountHandler interface {
	AuthMiddleware(next httprouter.Handle) httprouter.Handle
}

type accountHandler struct {
	AccountUseCase usecase.AccountUseCase
}

func NewAccountHandler(a usecase.AccountUseCase) AccountHandler {
	return &accountHandler{
		AccountUseCase: a,
	}
}

func (a *accountHandler) AuthMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		fmt.Println("start")
		token := r.URL.Query().Get("token")
		if token == "" {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}
		ctx := r.Context()
		a.AccountUseCase.SetAccountToCtxByToken(token, &ctx)
		r = r.WithContext(ctx)

		fmt.Println("end")
		next(w, r, ps)
	}
}
