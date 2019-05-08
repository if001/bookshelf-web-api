package tables

import "time"

type AuthToken struct {
	BaseModel
	Token       string `gorm:"type:varchar(45);"`
	ExpireTime  time.Time `sql:"not null;type:date"`
}

func (AuthToken) TableName() string {
	return "auth_token"
}
