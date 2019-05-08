package infrastructure

import (
	"bookshelf-web-api/domain/model"
	"bookshelf-web-api/domain/repository"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type accountRepository struct {
	DB *gorm.DB
}

func NewAccountRepository(db *gorm.DB) repository.AccountRepository {
	return &accountRepository{ DB : db }
}


func (c *accountRepository) Get(token string) (*[]model.Account, error) {
	var account []model.Account
	var authToken model.AuthToken

	err := c.DB.Joins("JOIN auth_token ON auth_token.account_id = accounts.id").
		Where("token = ?",token).
		Find(&account).
		Error
	if err == nil {
		err = c.DB.Where("token = ?", token).Find(&authToken).Error
		if err == nil {
			if authToken.ExpireTime.After(time.Now()) {
				// authToken.ExpireTime = time.Now().AddDate(0,3,0) // 3ヶ月伸ばす
				// err = c.DB.Save(&authToken).Error
				fmt.Println("expire time ok")
			} else {
				err = errors.New("expire time")
			}
		}
	}
	return &account, err
}
