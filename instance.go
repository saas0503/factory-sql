package sql

import "gorm.io/gorm"

func (db *DB) GetInstance() *gorm.DB {
	return db.Connect("")
}

func (db *DB) GetInstanceTenant(tenant string) *gorm.DB {
	return db.Connect(tenant)
}
