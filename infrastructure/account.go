package infrastructure

import (
	"bookshelf-web-api/domain/model"
	"bookshelf-web-api/domain/repository"
	"bookshelf-web-api/domain/service"
	"github.com/jinzhu/gorm"
)

type accountRepository struct {
	DB *gorm.DB
}

func NewAccountRepository(db *gorm.DB) repository.AccountRepository {
	return &accountRepository{ DB : db }
}

var account model.Account

func (c *accountRepository) Get(token string) (*model.Account, service.RecodeNotFoundError) {
	err := c.DB.Joins("JOIN auth_token ON auth_token.account_id = accounts.id").
		Where("token = ?",token).
		Find(&account).
		Error
	return &account, err
}
