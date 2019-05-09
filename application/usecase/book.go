package usecase

import (
	"bookshelf-web-api/domain/model"
	"bookshelf-web-api/domain/repository"
	"bookshelf-web-api/domain/service"
	"bookshelf-web-api/application/usecase/form"
	"encoding/json"
)

type BookUseCase interface {
	BookListUseCase(account model.Account) (*[]model.Book, service.RecodeNotFoundError)
	//BookFindUseCase(id int64, account model.Account) (*[]model.Book, service.RecodeNotFoundError)
	CreateBook(bookRequest model.Book, account model.Account) (*model.Book, service.RecodeNotFoundError)
	//UpdateBook(id int64, bookRequest model.BookRequest, account model.Account) (*model.Book, service.RecodeNotFoundError)
	//DescriptionFindUseCase(id int64) (*[]model.Description, service.RecodeNotFoundError)
	//DescriptionCreateUseCase(id int64, description string) (*model.Description, service.RecodeNotFoundError)
	BookRequestBind(body []byte) (*model.Book, error)
}

type bookUseCase struct {
	BookR repository.BookRepository
	AuthorR repository.AuthorRepository
	CategoryR repository.CategoryRepository
}

func NewBookUseCase(bookR repository.BookRepository, authorR repository.AuthorRepository, categoryR repository.CategoryRepository) BookUseCase {
	return &bookUseCase{
		BookR: bookR,
		AuthorR: authorR,
		CategoryR:categoryR,
	}
}

func (u *bookUseCase) BookListUseCase(account model.Account) (*[]model.Book, service.RecodeNotFoundError) {
	books, err := u.BookR.GetBooks(account.ID)
	return books, service.RecodeNotFoundError(err)
}

func (u *bookUseCase) CreateBook(book model.Book, account model.Account) (*model.Book, service.RecodeNotFoundError) {
	newBook, err := u.BookR.CreateBook(book, account)
	return newBook, service.RecodeNotFoundError(err)
}


func (u *bookUseCase) BookRequestBind(body []byte) (*model.Book, error) {
	var bookRequest form.BookRequest
	err := json.Unmarshal(body, &bookRequest)
	if err != nil {
		return nil, err
	}
	book := model.Book{}
	book.Name = bookRequest.Title

	author := &model.Author{}
	if bookRequest.Author == "" {
		author = nil
	} else {
		author.Name = bookRequest.Author
	}
	isExist, err := u.AuthorR.IsExistAuthor(author)
	if err != nil {
		return nil, err
	}

	if isExist {
		book.Author = nil
	} else {
		book.Author = author
	}

	book.NextBookID = bookRequest.NextBookId
	book.PrevBookID = bookRequest.PrevBookId

	categories := []model.Category{}
	for i := range bookRequest.Categories {
		category := model.Category{}
		category.Name = bookRequest.Categories[i]
		categories = append(
			categories,
			category,
		)
	}

	notExistCategories,err := u.CategoryR.GetNotExistCategories(categories)
	if err != nil {
		return nil, err
	}
	book.Categories = *notExistCategories
	return &book, nil
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
//func (u *bookUseCase) UpdateBook(id int64, bookRequest model.BookRequest, account model.Account) (*model.Book, service.RecodeNotFoundError) {
//	book, err := u.BookRepo.UpdateBook(id, bookRequest, account)
//	return book, service.RecodeNotFoundError(err)
//}
