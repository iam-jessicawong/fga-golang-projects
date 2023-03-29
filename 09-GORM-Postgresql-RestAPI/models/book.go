package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Book struct {
	// *gorm.Model
	ID          int    `gorm:"primaryKey" json:"id"`
	Title       string `gorm:"not null;type:varchar(100)" json:"title" binding:"required"`
	Author      string `gorm:"not null;type:varchar(100)" json:"author" binding:"required"`
	Description string `gorm:"not null;type:text" json:"description" binding:"required"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// method || hooks
func (b *Book) BeforeCreate(tx *gorm.DB) (err error) {
	if len(b.Title) < 2 {
		err = errors.New("title is too short")
	}
	if len(b.Description) < 5 {
		err = errors.New("description is too short")
	}
	if len(b.Author) < 2 {
		err = errors.New("author is too short")
	}

	return
}

/* func (b *Book) BeforeUpdate(tx *gorm.DB) (err error) {
	if len(b.Title) < 2 {
		err = errors.New("title is too short")
	}
	if len(b.Description) < 5 {
		err = errors.New("description is too short")
	}
	if len(b.Author) < 2 {
		err = errors.New("author is too short")
	}

	return
} */
