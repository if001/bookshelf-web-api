package model

import "time"

type BaseModel struct {
	ID        int64 `gorm:"primary_key" sql:"AUTO_INCREMENT"`
	CreatedAt time.Time `gorm;"default:CURRENT_TIMESTAMP" sql:"not null;type:date"`
	UpdatedAt time.Time `gorm;"default:CURRENT_TIMESTAMP" sql:"not null;type:date"`
}
