package usecase

import (
	"bookshelf-web-api/domain/model"
	"bookshelf-web-api/domain/repository"
	"encoding/json"
	"io"
	"bookshelf-web-api/application/usecase/form"
)

type BookUseCase interface {
	BookListUseCase(account model.Account) (*[]model.Book, error)
	CreateBook(book model.Book, account model.Account) (*model.Book, error)
	BookFindUseCase(id int64, account model.Account) (*model.Book, error)
	UpdateBook(book model.Book, account model.Account) (*model.Book, error)
	//DescriptionFindUseCase(id int64) (*[]model.Description, service.RecodeNotFoundError)
	//DescriptionCreateUseCase(id int64, description string) (*model.Description, service.RecodeNotFoundError)
	BookRequestBind(body io.ReadCloser) (*model.Book, error)
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

func (u *bookUseCase) BookListUseCase(account model.Account) (*[]model.Book, error) {
	books, err := u.BookR.GetBooks(account.ID)
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (u *bookUseCase) CreateBook(book model.Book, account model.Account) (*model.Book, error) {
	newBook, err := u.BookR.CreateBook(book, account)
	if err != nil {
		return nil, err
	}
	return newBook, nil
}


func (u *bookUseCase) BookRequestBind(body io.ReadCloser) (*model.Book, error) {
	var bookRequest form.BookRequest
	// TODO 存在しないkeyがrequestにあったらbad requestにしたい
	err := json.NewDecoder(body).Decode(&bookRequest)
	if err != nil {
		return nil, err
	}
	book := model.Book{}
	book.Name = bookRequest.Title
	author := &model.Author{}
	if bookRequest.AuthorId == 0 {
		author = nil
	} else {
		author, err = u.AuthorR.GetAuthor(bookRequest.AuthorId)
		if err != nil {
			return nil, err
		}
	}
	book.Author = author

	book.NextBookID = bookRequest.NextBookId
	book.PrevBookID = bookRequest.PrevBookId
	categories,err := u.CategoryR.GetByIds(bookRequest.CategoryIds)
	if err != nil {
		return nil, err
	}
	book.Categories = *categories
	return &book, nil
}

func (u *bookUseCase) BookFindUseCase(id int64, account model.Account) (*model.Book, error) {
	book, err := u.BookR.FindBook(id, account)

	isExist, err := u.AuthorR.IsExistAuthor(book.Author)
	if err != nil {
		return nil, err
	}
	if isExist {
		book.Author = nil
	}
	notExistCategories,err := u.CategoryR.GetNotExistCategories(book.Categories)
	book.Categories = *notExistCategories

	if err != nil {
		return nil, err
	}
	return book, nil
}

func (u *bookUseCase) UpdateBook(book model.Book, account model.Account) (*model.Book, error) {
	newBook, err := u.BookR.UpdateBook(book, account)
	if err != nil {
		return nil, err
	}
	return newBook, nil
}


//
//func (u *bookUseCase) DescriptionCreateUseCase(id int64, description string) (*model.Description, service.RecodeNotFoundError) {
//	newDescription, err := u.BookRepo.CreateDescription(id, description)
//	return newDescription, service.RecodeNotFoundError(err)
//}
//
