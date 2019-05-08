package usecase

import (
	"bookshelf-web-api/domain/model"
	"bookshelf-web-api/domain/repository"
	"bookshelf-web-api/domain/service"
	"context"
	"fmt"
)

type BookUseCase interface {
	BookListUseCase(ctx context.Context) (*[]model.Book, service.RecodeNotFoundError)
	//BookFindUseCase(id int64, account model.Account) (*[]model.Book, service.RecodeNotFoundError)
	//DescriptionFindUseCase(id int64) (*[]model.Description, service.RecodeNotFoundError)
	//DescriptionCreateUseCase(id int64, description string) (*model.Description, service.RecodeNotFoundError)
	//CreateBook(bookRequest model.BookRequest, account model.Account) (*model.Book, service.RecodeNotFoundError)
	//UpdateBook(id int64, bookRequest model.BookRequest, account model.Account) (*model.Book, service.RecodeNotFoundError)
}

type bookUseCase struct {
	BookR repository.BookRepository
	AuthorR repository.AuthorRepository
	AccountR repository.AccountRepository
}

func NewBookUseCase(bookR repository.BookRepository, authorR repository.AuthorRepository, accountR repository.AccountRepository) BookUseCase {
	return &bookUseCase{
		BookR: bookR,
		AuthorR: authorR,
		AccountR: accountR,
	}
}

func (u *bookUseCase) BookListUseCase(ctx context.Context) (*[]model.Book, service.RecodeNotFoundError) {
	fmt.Println("use case")
	fmt.Println("use case", ctx.Value("account"))
	account, err  := u.AccountR.GetAccount(ctx)

	if err != nil {
		return nil, err
	}
	books, err := u.BookR.GetBooks(account.ID)
	return books, service.RecodeNotFoundError(err)
}


//func (u *bookUseCase) BookFindUseCase(id int64, account model.Account) (*[]model.Book, service.RecodeNotFoundError) {
//	book, err := u.BookRepo.FindBook(id, account)
//	return book, service.RecodeNotFoundError(err)
//}
//func (u *bookUseCase) DescriptionFindUseCase(id int64) (*[]model.Description, service.RecodeNotFoundError) {
//	book, err := u.BookRepo.FindDescriptions(id)
//	return book, service.RecodeNotFoundError(err)
//}
//
//func (u *bookUseCase) DescriptionCreateUseCase(id int64, description string) (*model.Description, service.RecodeNotFoundError) {
//	newDescription, err := u.BookRepo.CreateDescription(id, description)
//	return newDescription, service.RecodeNotFoundError(err)
//}
//
//func (u *bookUseCase) CreateBook(bookRequest model.BookRequest, account model.Account) (*model.Book, service.RecodeNotFoundError) {
//	book, err := u.BookRepo.CreateBook(bookRequest, account)
//	return book, service.RecodeNotFoundError(err)
//}
//func (u *bookUseCase) UpdateBook(id int64, bookRequest model.BookRequest, account model.Account) (*model.Book, service.RecodeNotFoundError) {
//	book, err := u.BookRepo.UpdateBook(id, bookRequest, account)
//	return book, service.RecodeNotFoundError(err)
//}
