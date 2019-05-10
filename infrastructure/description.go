package infrastructure

import (
	// "bookshelf-web-api/domain/model"
	"bookshelf-web-api/domain/repository"
	// "bookshelf-web-api/domain/service"
	// "errors"
	"github.com/jinzhu/gorm"
	"bookshelf-web-api/domain/model"
	"bookshelf-web-api/infrastructure/tables"
)

type descriptionRepository struct {
	DB *gorm.DB
}

func NewDescriptionRepository(db *gorm.DB) repository.DescriptionRepository {
	return &descriptionRepository{ DB : db }
}



func (c *descriptionRepository) FindDescriptions(id int64) (*[]model.Description, error) {
	descriptions := []model.Description{}
	descriptionsTable := []tables.Description{}

	err := c.DB.Where("book_id = ?", id).Find(&descriptionsTable).Error
	if err != nil {
		return nil, err
	}
	for i := range descriptionsTable {
		description := model.Description{}
		description.Fill(
			descriptionsTable[i].ID,
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
