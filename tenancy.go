package sql

import (
	"log"

	"gorm.io/gorm"
)

type Tenancy struct {
	Name string
	DB   *gorm.DB
}

func (tenancy *Tenancy) InitModels(models ...interface{}) {
	tenancy.DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")

	log.Println("Running Migrations")
	err := tenancy.DB.AutoMigrate(models...)

	if err != nil {
		panic(err)
	}
}
