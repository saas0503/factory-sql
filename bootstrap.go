package sql

import "github.com/saas0503/factory-sql/tenant"

func Bootstrap() {
	dbMaster := Connect("")

	var tenants []*tenant.Tenant

	result := dbMaster.Db.Find(&tenants)
	if result.Error != nil {
		panic(result.Error.Error())
	}

	for _, tn := range tenants {
		Connect(tn.Code)
	}
}
