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
	bookR := infrastructure.NewBookRepository(db)
	accountR := infrastructure.NewAccountRepository(db)
	authorR := infrastructure.NewAuthorRepository(db)
	categoryR := infrastructure.NewCategoryRepository(db)
	dr := infrastructure.NewDescriptionRepository(db)

	categoryUseCase := usecase.NewCategoryUseCase(cr)
	bookUseCase := usecase.NewBookUseCase(bookR, authorR, categoryR)
	accountUseCase := usecase.NewAccountUseCase(accountR)
	descriptionUseCase := usecase.NewDescriptionUseCase(dr)

	uh := handler.NewCategoryHandler(categoryUseCase)
	bh := handler.NewBookHandler(bookUseCase, accountUseCase)
	ah := handler.NewAccountHandler(accountUseCase)
	dh := handler.NewDescriptionHandler(descriptionUseCase)

	api := handler.NewApiHandler(uh,bh,ah,dh)
	
	r := router.Route(api)

	fmt.Printf("[START] server. port: %s\n", addr)

	if err := http.ListenAndServe(addr, handler.Log(r)); err != nil {
		panic(fmt.Errorf("[FAILED] start sever. err: %v", err))
	}
}
