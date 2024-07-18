package sql

import (
	"log"

	"gorm.io/gorm"
)

type Tenancy struct {
	Name string
	Db   *gorm.DB
}

func (tenancy *Tenancy) InitModels(models ...interface{}) {
	log.Println("Running Migrations")
	err := tenancy.Db.AutoMigrate(models...)

	if err != nil {
		panic(err)
	}
}
