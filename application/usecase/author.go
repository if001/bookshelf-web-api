package usecase
//
//import (
//	"bookshelf-web-api/domain/model"
//	"bookshelf-web-api/domain/repository"
//	"bookshelf-web-api/domain/service"
//)
//
//type AuthorUseCase interface {
//	CreateBook(author string) (*[]model.Author, service.RecodeNotFoundError)
//}
//
//type authorUseCase struct {
//	AuthorRepo repository.AuthorRepository
//}
//
//func NewAuthorUseCase(cr repository.AuthorRepository) AuthorUseCase {
//	return &authorUseCase{
//		AuthorRepo: cr,
//	}
//}
//
//func (u *authorUseCase) CreateBook(author string) (*[]model.Author, service.RecodeNotFoundError) {
//	a, err := u.AuthorRepo.CreateBook(author)
//	return a, service.RecodeNotFoundError(err)
//}
