package sql

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var ConnectionMap = make(map[string]*gorm.DB)

type ConnectionOptions struct {
	Host string
	Port string
	User string
	Pass string
}

func Connect(tenant string) *Tenancy {
	opts := ConnectionOptions{
		Host: os.Getenv("DB_HOST"),
		Port: os.Getenv("DB_PORT"),
		User: os.Getenv("DB_USER"),
		Pass: os.Getenv("DB_PASS"),
	}
	if ConnectionMap[tenant] != nil {
		dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", opts.Host, opts.Port, opts.User, opts.Pass, tenant)

		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		ConnectionMap[tenant] = db
	}
	return &Tenancy{
		Name: tenant,
		Db:   ConnectionMap[tenant],
	}
}
