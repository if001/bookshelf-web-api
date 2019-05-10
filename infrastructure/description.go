package infrastructure

import (
	// "bookshelf-web-api/domain/model"
	"bookshelf-web-api/domain/repository"
	// "bookshelf-web-api/domain/service"
	// "errors"
	"github.com/jinzhu/gorm"
	"bookshelf-web-api/domain/model"
	"bookshelf-web-api/infrastructure/tables"
	"errors"
)

type descriptionRepository struct {
	DB *gorm.DB
}

func NewDescriptionRepository(db *gorm.DB) repository.DescriptionRepository {
	return &descriptionRepository{ DB : db }
}



func (r *descriptionRepository) FindDescriptions(id int64) (*[]model.Description, error) {
	descriptions := []model.Description{}
	descriptionsTable := []tables.Description{}

	err := r.DB.Where("book_id = ?", id).Find(&descriptionsTable).Error
	if err != nil {
		return nil, err
	}
	for i := range descriptionsTable {
		description := model.Description{}
		description.Fill(
			descriptionsTable[i].ID,
			id,
			descriptionsTable[i].Description,
			descriptionsTable[i].CreatedAt,
			descriptionsTable[i].UpdatedAt,
		)
		descriptions = append(
			descriptions,
			description,
		)
	}
	return &descriptions, err
}


func (r *descriptionRepository) CreateDescription(description model.Description) (*model.Description, error) {
	var books []model.Book
	err := r.DB.Where("id = ?", description.BookId).Find(&books).Error
	if err != nil {
		return nil, err
	}
	if len(books) == 0 {
		return nil, errors.New("record not found")
	}

	descriptionTable := tables.Description{}
	descriptionTable.BookId = description.BookId
	descriptionTable.Description = description.Content

	err = r.DB.Create(&descriptionTable).Error
	if err != nil {
		return nil, err
	}
	newDescription := model.Description{}

	newDescription.Fill(
		descriptionTable.ID,
		description.BookId,
		description.Content,
		description.CreatedAt,
		description.UpdatedAt,
	)
	return &newDescription, err
}



//func (c *descriptionRepository) Update(id int64, descriptionRequest model.DescriptionRequest) (*model.Description, service.RecodeNotFoundError) {
//	var descriptionModelForBind []model.Description
//	err := c.DB.Where("id = ?", id).Find(&descriptionModelForBind).Error
//	if err != nil {
//		return nil, err
//	}
//	if len(descriptionModelForBind) == 0 {
//		return nil, errors.New("not found")
//	}
//	descriptionModel := descriptionModelForBind[0]
//
//	descriptionModel.Description = descriptionRequest.Description
//	err = c.DB.Save(&descriptionModel).Error
//	if err != nil {
//		return nil, err
//	}
//	return &descriptionModel, err
//}
