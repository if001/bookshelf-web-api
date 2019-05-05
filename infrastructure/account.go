package infrastructure

import (
	"bookshelf-web-api/domain/model"
	"bookshelf-web-api/domain/repository"
	"errors"
	"github.com/jinzhu/gorm"
	"time"
)

type accountRepository struct {
	DB *gorm.DB
}

func NewAccountRepository(db *gorm.DB) repository.AccountRepository {
	return &accountRepository{ DB : db }
}

var account model.Account
var authToken model.AuthToken

func (c *accountRepository) Get(token string) (*model.Account, error) {
	err := c.DB.Joins("JOIN auth_token ON auth_token.account_id = accounts.id").
		Where("token = ?",token).
		Find(&account).
		Error
	if err == nil {
		err = c.DB.Where("token = ?", token).Find(&authToken).Error
		if err == nil {
			if authToken.ExpireTime.Before(time.Now()) {
				authToken.ExpireTime = time.Now()
				err = c.DB.Save(&authToken).Error
			} else {
				err = errors.New("expire time")
			}
		}
	}
	return &account, err
}
