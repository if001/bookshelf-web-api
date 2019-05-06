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
	ar := infrastructure.NewAccountRepository(db)
	ar2 := infrastructure.NewAuthorRepository(db)

	categoryUseCase := usecase.NewCategoryUseCase(cr)
	bookUseCase := usecase.NewBookUseCase(br, ar2)
	accountUseCase := usecase.NewAccountUseCase(ar)

	uh := handler.NewCategoryHandler(categoryUseCase)
	bh := handler.NewBookHandler(bookUseCase)
	ah := handler.NewAccountHandler(accountUseCase)

	api := handler.NewApiHandler(uh,bh,ah)
	
	r := router.Route(api)

	fmt.Printf("[START] server. port: %s\n", addr)

	if err := http.ListenAndServe(addr, handler.Log(r)); err != nil {
		panic(fmt.Errorf("[FAILED] start sever. err: %v", err))
	}
}
