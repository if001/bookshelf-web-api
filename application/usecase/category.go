package usecase

import (
	"bookshelf-web-api/domain/repository"
	"bookshelf-web-api/domain/model"
)

type CategoryUseCase interface {
	CategoryUseCase() (*model.Category, error)
}

type categoryUseCase struct {
	CategoryRepo repository.CategoryRepository
}

func NewCategoryUseCase(cr repository.CategoryRepository) CategoryUseCase {
	return &categoryUseCase{
		CategoryRepo: cr,
	}
}

func (u *categoryUseCase) CategoryUseCase() (*model.Category, error) {
	cate, err := u.CategoryRepo.Get()
	return cate, err
}
