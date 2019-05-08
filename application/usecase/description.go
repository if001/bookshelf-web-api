package usecase

import (
	"bookshelf-web-api/domain/repository"
)

type DescriptionUseCase interface {
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

//func (u *descriptionUseCase) DescriptionUpdateUseCase(id int64, description model.DescriptionRequest) (*model.Description, error) {
//	d, err := u.DescriptionRepo.Update(id, description)
//	return d, err
//}
