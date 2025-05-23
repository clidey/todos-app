package models

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Title     string         `json:"title"`
	Completed bool           `json:"completed"`
	Order     int            `json:"order" gorm:"default:0"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
