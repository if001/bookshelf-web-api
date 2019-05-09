package infrastructure

import (
	"bookshelf-web-api/domain/repository"
	"github.com/jinzhu/gorm"
	"bookshelf-web-api/domain/model"
	"bookshelf-web-api/infrastructure/tables"
	"errors"
)

type authorRepository struct {
	DB *gorm.DB
}

func NewAuthorRepository(db *gorm.DB) repository.AuthorRepository {
	return &authorRepository{ DB : db }
}


func (r *authorRepository) IsExistAuthor(author *model.Author) (bool, error) {
	authorTable := []tables.Author{}
	if author == nil {
		return true, nil
	}
	err := r.DB.Where("name = ?", author.Name).Find(&authorTable).Error
	if err != nil {
		return false, err
	}
	if len(authorTable) == 0 {
		return false, nil
	} else {
		return true, nil
	}
}

func (r *authorRepository) GetAuthor(id int64) (*model.Author, error) {
	authorTable := []tables.Author{}
	err := r.DB.Where("id = ?",id).Find(&authorTable).Error
	if err != nil {
		return nil, err
	}
	if len(authorTable) == 0 {
		return nil, errors.New("record not found")
	}

	author := model.Author{}
	author.Fill(
		authorTable[0].ID,
		authorTable[0].Name,
		authorTable[0].CreatedAt,
		authorTable[0].UpdatedAt,
	)
	return &author, nil
}