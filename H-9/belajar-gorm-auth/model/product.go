package model

import (
	"database/sql"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model

	Name        string       `json:"name"`
	Description string       `json:"description"`
	Price       float32      `json:"price"`
	Stock       float32      `json:"stock"`
	Expired     sql.NullTime `json:"expired,omitempty"`
}
