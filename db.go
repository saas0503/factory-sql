package sql

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

type DB struct {
	isTenancy     bool
	ConnectionMap map[string]*gorm.DB
}

type ConnectionOptions struct {
	Host string
	Port string
	User string
	Pass string
}

func NewDB(isTenancy bool) *DB {
	return &DB{
		isTenancy:     isTenancy,
		ConnectionMap: make(map[string]*gorm.DB),
	}
}

func (db *DB) Connect(tenant string) *gorm.DB {
	opts := ConnectionOptions{
		Host: os.Getenv("DB_HOST"),
		Port: os.Getenv("DB_PORT"),
		User: os.Getenv("DB_USER"),
		Pass: os.Getenv("DB_PASS"),
	}

	if !db.isTenancy || tenant == "" {
		tenant = os.Getenv("DB_NAME")
	}

	if db.ConnectionMap[tenant] == nil {
		dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", opts.Host, opts.Port, opts.User, opts.Pass, tenant)

		conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		db.ConnectionMap[tenant] = conn
	}
	return db.ConnectionMap[tenant]
}

func (db *DB) InitModels(models ...interface{}) {
	name := os.Getenv("DB_NAME")

	migration(db.ConnectionMap[name], models...)
}

func (db *DB) InitModelsTenant(tenant string, models ...interface{}) {
	migration(db.ConnectionMap[tenant], models...)
}

func (db *DB) Disconnect(tenant string) {
	delete(db.ConnectionMap, tenant)
}

func migration(db *gorm.DB, models ...interface{}) {
	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")

	log.Println("Running Migrations")
	err := db.AutoMigrate(models...)

	if err != nil {
		panic(err)
	}
}
