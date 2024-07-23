package sql

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type DB struct {
	opt           ConnectionOptions
	ConnectionMap map[string]*gorm.DB
}

type ConnectionOptions struct {
	Host      string
	Port      string
	Username  string
	Password  string
	Database  string
	IsTenancy bool
}

func NewDB(conn ConnectionOptions) *DB {
	return &DB{
		opt:           conn,
		ConnectionMap: make(map[string]*gorm.DB),
	}
}

func (db *DB) Connect(tenant string) *gorm.DB {
	opts := db.opt
	if !opts.IsTenancy || tenant == "" {
		tenant = db.opt.Database
	}

	if db.ConnectionMap[tenant] == nil {
		dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", opts.Host, opts.Port, opts.Username, opts.Password, tenant)

		conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		db.ConnectionMap[tenant] = conn
	}
	return db.ConnectionMap[tenant]
}

func (db *DB) InitModels(models ...interface{}) {
	name := db.opt.Database

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
