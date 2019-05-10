package usecase

import (
	"bookshelf-web-api/domain/repository"
	"bookshelf-web-api/domain/model"
)

type DescriptionUseCase interface {
	DescriptionFindUseCase(id int64) (*[]model.Description, error)
	// DescriptionUpdateUseCase(id int64, description model.DescriptionRequest) (*model.Description, error)
}

type descriptionUseCase struct {
	DescriptionRepo repository.DescriptionRepository
}

func NewDescriptionUseCase(r repository.DescriptionRepository) DescriptionUseCase {
	return &descriptionUseCase{
		DescriptionRepo: r,
	}
}

func (u *descriptionUseCase) DescriptionFindUseCase(id int64) (*[]model.Description, error) {
	book, err := u.DescriptionRepo.FindDescriptions(id)
	if err != nil {
		return nil, err
	}
	return book, nil
}
//func (u *descriptionUseCase) DescriptionUpdateUseCase(id int64, description model.DescriptionRequest) (*model.Description, error) {
//	d, err := u.DescriptionRepo.Update(id, description)
//	return d, err
//}
