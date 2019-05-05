package model


type Account struct {
	BaseModel
	UUID string `gorm:"type:varchar(40);"`
}

func (Account) TableName() string {
	return "books"
}
