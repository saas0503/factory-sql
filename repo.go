package sql

type Repo[M any] struct {
	Name string
	DB   *Tenancy
}

func NewRepo[M any](name string, db *Tenancy) *Repo[M] {
	return &Repo[M]{
		Name: name,
		DB:   db,
	}
}
