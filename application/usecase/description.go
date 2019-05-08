package usecase

import (
	"bookshelf-web-api/domain/model"
	"bookshelf-web-api/domain/repository"
)

type DescriptionUseCase interface {
	DescriptionUpdateUseCase(id int64, description string) (*[]model.Description, error)
}

type descriptionUseCase struct {
	DescriptionRepo repository.DescriptionRepository
}

func NewDescriptionUseCase(r repository.DescriptionRepository) DescriptionUseCase {
	return &descriptionUseCase{
		DescriptionRepo: r,
	}
}

func (u *descriptionUseCase) DescriptionUpdateUseCase(id int64, description string) (*[]model.Description, error) {
	d, err := u.DescriptionRepo.Update(id, description)
	return d, err
}
