package infrastructure

import (
	"bookshelf-web-api/domain/repository"
	"errors"
	"github.com/jinzhu/gorm"
	"context"
	"bookshelf-web-api/infrastructure/tables"
	"time"
	"fmt"
)

type accountRepository struct {
	DB *gorm.DB
}

func NewAccountRepository(db *gorm.DB) repository.AccountRepository {
	return &accountRepository{ DB : db }
}

func (c *accountRepository) GetAccount(ctx context.Context) (*tables.Account, error) {
	account,ok := ctx.Value("account").(tables.Account)
	if ok {
		return &account, nil
	} else {
		return nil, errors.New("bind error")
	}
}


func (c *accountRepository) SetAccount(token string, ctx *context.Context) (error) {
	var account []tables.Account
	var authToken tables.AuthToken

	err := c.DB.Joins("JOIN auth_token ON auth_token.account_id = accounts.id").
		Where("token = ?",token).
		Find(&account).
		Error
	if err != nil {
		return err
	}
	err = c.DB.Where("token = ?", token).Find(&authToken).Error
	if err != nil {
		return err
	}

	if authToken.ExpireTime.After(time.Now()) {
		// authToken.ExpireTime = time.Now().AddDate(0,3,0) // 3ヶ月伸ばす
		// err = c.DB.Save(&authToken).Error
		fmt.Println("expire time ok")
	} else {
		return errors.New("expire time")
	}
	if len(account) == 0 {
		return errors.New("record not found")
	}
	*ctx = context.WithValue(*ctx, "account", account[0])
	return nil
}
