package sql

import (
	"github.com/saas0503/factory-sql/tenant"
)

func Bootstrap() {
	dbMaster := Connect("")
	dbMaster.InitModels(&tenant.Tenant{})

	var tenants []*tenant.Tenant

	result := dbMaster.DB.Find(&tenants)
	if result.Error != nil {
		panic(result.Error.Error())
	}

	for _, tn := range tenants {
		Connect(tn.Code)
	}
}
