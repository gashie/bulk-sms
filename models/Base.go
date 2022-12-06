package models

import (
	"time"

)

type Base struct {
	// ID        uuid.UUID  `gorm:"type:char(36);primaryKey"`
	CreatedAt time.Time  `gorm:"type:timestamp;" json:"-"`
	UpdatedAt time.Time  `gorm:"type:Date;" json:"-"`
	DeletedAt *time.Time `gorm:"type:Date;" json:"-"`
}
