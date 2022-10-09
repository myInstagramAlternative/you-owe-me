package repository

import (
	"log"
	"you-owe-me/model"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type oweRepository struct {
	DB *gorm.DB
}

// OweRepository : represent the Owe's repository contract
type OweRepository interface {
	GetOwe(uuid.UUID) (model.Owe, error)

	Migrate() error
}

// NewOweRepository -> returns new owe repository
func NewOweRepository(db *gorm.DB) OweRepository {
	return oweRepository{
		DB: db,
	}
}

func (o oweRepository) Migrate() error {
	log.Print("[OweRepository]...Migrate")
	return o.DB.AutoMigrate(&model.Owe{})
}

func (o oweRepository) GetOwe(uuid uuid.UUID) (owe model.Owe, err error) {
	return owe, o.DB.First(&owe, uuid).Error
}
