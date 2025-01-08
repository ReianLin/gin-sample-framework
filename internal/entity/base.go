package entity

import (
	"time"

	"gorm.io/gorm"
)

type BaseEntity struct {
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;type:timestamp;not null;default:now()"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;type:timestamp;not null;default:now()"`
}

func (b *BaseEntity) BeforeCreate(tx *gorm.DB) error {
	b.CreatedAt = time.Now().UTC()
	b.UpdatedAt = time.Now().UTC()
	return nil
}

func (b *BaseEntity) BeforeUpdate(tx *gorm.DB) error {
	b.UpdatedAt = time.Now().UTC()
	return nil
}
