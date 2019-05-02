// mainHTTPSample.go
package main

import (
	"fmt"
	"net/http"
	"bookshelf-web-api/interfaces/api/handler"
	"bookshelf-web-api/interfaces/api/router"
	"bookshelf-web-api/application/usecase"
	"bookshelf-web-api/infrastructure"
	"bookshelf-web-api/infrastructure/mysql"
)

func main() {
	var addr =  ":8181"

	db := mysql.GetDBConn()
	cr := infrastructure.NewCategoryRepository(db)
	br := infrastructure.NewBookRepository(db)

	categoryUseCase := usecase.NewCategoryUseCase(cr)
	bookUseCase := usecase.NewBookUseCase(br)

	uh := handler.NewCategoryHandler(categoryUseCase)
	bh := handler.NewBookHandler(bookUseCase)
	ah := handler.NewApiHandler(uh,bh)

	router := router.Route(ah)

	fmt.Printf("[START] server. port: %s\n", addr)

	if err := http.ListenAndServe(addr, handler.Log(router)); err != nil {
		panic(fmt.Errorf("[FAILED] start sever. err: %v", err))
	}
}