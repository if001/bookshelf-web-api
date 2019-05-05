package model

import "time"

type BaseModel struct {
	ID        int64 `gorm:"primary_key" sql:"AUTO_INCREMENT"`
	CreatedAt time.Time `sql:"not null;type:date"`
	UpdatedAt time.Time `sql:"not null;type:date"`
}
