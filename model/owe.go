package model

import (
	"time"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

// User -> model for users table
type Base struct {
	ID        uuid.UUID `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (base *Base) BeforeCreate(tx *gorm.DB) error {
	uuid, err := uuid.NewV4()
	if err != nil {
		return err
	}
	base.ID = uuid
	return nil
}

type Owe struct {
	Base
	Currency   string `json:"currency" default:"RSD"`
	Ammount    string `json:"ammount"`
	From       string `json:"from"`
	To         string `json:"to"`
	FromSigned bool   `json:"from_signed default:false"`
	ToSigned   bool   `json:"to_signed default:false"`
}

// TableName --> Table for Product Model
func (Owe) TableName() string {
	return "owes"
}
