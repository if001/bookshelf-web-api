package handler

import (
	"bookshelf-web-api/application/usecase"
	"context"
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
		} else {
			account, err := a.AccountUseCase.AccountGetUseCase(token)
			if err != nil {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			} else {
				ctx := context.WithValue(r.Context(),"account",account)
				r = r.WithContext(ctx)
				fmt.Println("end")
				next(w, r, ps)
			}
		}

	}
}
