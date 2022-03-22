package model

import "gorm.io/gorm"

//User -> model for users table
type Log struct {
	gorm.Model
	Host    string  `json:"host"`
	Action  string  `json:"action" `
	Details Details `gorm:"embedded;embeddedPrefix:details_"`
	Status  string  `json:"status"`
}

type Details struct {
	Type string
	Data string
}

//TableName --> Table for Product Model
func (Log) TableName() string {
	return "logs"
}
