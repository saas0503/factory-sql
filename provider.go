package sql

import (
	"net/http"
)

type Provider struct {
	DB     *Tenancy
	Tenant string
}

func NewProvider(r *http.Request) *Provider {
	tenant := r.Context().Value(TenantToken).(string)
	tenancyDB := Connect(tenant)
	return &Provider{
		Tenant: tenant,
		DB:     tenancyDB,
	}
}
