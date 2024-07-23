package sql

import (
	"github.com/saas0503/factory-sql/tenant"
)

func (db *DB) Bootstrap() {
	dbMaster := db.Connect("")
	db.InitModels(&tenant.Tenant{})

	var tenants []*tenant.Tenant

	result := dbMaster.Find(&tenants)
	if result.Error != nil {
		panic(result.Error.Error())
	}

	for _, tn := range tenants {
		db.Connect(tn.Code)
	}
}
