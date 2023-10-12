package db

import (
	"app/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Con *gorm.DB

func Connect() {
	db, err := gorm.Open(postgres.Open("user=postgres password=SN5h3wqBp0qucUay host=db.rpgbhjzanthywqkwuyed.supabase.co port=5432 dbname=postgres"))
	if err != nil {
		panic(err)
	}

	Con = db

	err = Con.AutoMigrate(
		&model.Product{},
	)
	if err != nil {
		panic(err)
	}

}
