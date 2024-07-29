package fsql

import (
	"context"
	"net/http"
)

type Token string

const TenantToken Token = "tenant"

func TenantMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tenant := r.Header.Get("X-API-KEY")

		if tenant != "" {
			ctx := context.WithValue(r.Context(), TenantToken, tenant)
			r = r.WithContext(ctx)
		}

		next.ServeHTTP(w, r)
	})
}
