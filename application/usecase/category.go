package usecase

import (
	"bookshelf-web-api/domain/repository"
	"bookshelf-web-api/domain/model"
)

type CategoryUseCase interface {
	CategoryUseCase() (*[]model.Category, error)
	CategoryLogicalDeleteCase(bookId int64, categoryId int64) (error)
}

type categoryUseCase struct {
	CategoryRepo repository.CategoryRepository
}

func NewCategoryUseCase(cr repository.CategoryRepository) CategoryUseCase {
	return &categoryUseCase{
		CategoryRepo: cr,
	}
}

func (u *categoryUseCase) CategoryUseCase() (*[]model.Category, error) {
	categories, err := u.CategoryRepo.GetCategories()
	if err != nil {
		return nil, err
	}
	return categories, nil
}
func (u *categoryUseCase) CategoryLogicalDeleteCase(bookId int64, categoryId int64) (error) {
	err := u.CategoryRepo.LogicalDelete(bookId, categoryId)
	if err != nil {
		return err
	}
	return nil
}