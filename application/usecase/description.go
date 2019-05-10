package usecase

import (
	"bookshelf-web-api/domain/repository"
	"bookshelf-web-api/domain/model"
	"io"
	"bookshelf-web-api/application/usecase/form"
	"encoding/json"
	"errors"
)

type DescriptionUseCase interface {
	DescriptionFindUseCase(id int64) (*[]model.Description, error)
	DescriptionRequestBind(bookId int64, body io.ReadCloser) (*model.Description, error)
	DescriptionCreateUseCase(description model.Description) (*model.Description, error)
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


func (u *descriptionUseCase) DescriptionRequestBind(bookId int64, body io.ReadCloser) (*model.Description, error) {
	var descriptionRequest form.DescriptionRequest
	// TODO 存在しないkeyがrequestにあったらbad requestにしたい
	err := json.NewDecoder(body).Decode(&descriptionRequest)
	if err != nil {
		return nil, err
	}
	description := model.Description{}
	description.BookId = bookId
	description.Content = descriptionRequest.Description
	if description.BookId == 0 || description.Content == "" {
		return nil, errors.New("request bind error")
	}
	return &description, nil
}


func (u *descriptionUseCase) DescriptionCreateUseCase(description model.Description) (*model.Description, error) {
	newDescription, err := u.DescriptionRepo.CreateDescription(description)
	if err != nil {
		return nil, err
	}
	return newDescription, nil
}



//func (u *descriptionUseCase) DescriptionUpdateUseCase(id int64, description model.DescriptionRequest) (*model.Description, error) {
//	d, err := u.DescriptionRepo.Update(id, description)
//	return d, err
//}
