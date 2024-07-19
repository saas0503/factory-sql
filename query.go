package sql

import "gorm.io/gorm"

type QueryOptions struct {
	Where  interface{}
	Offset *int
	Limit  *int
	Order  interface{}
	Select []string
}

func (tenancy *Tenancy) Find(model interface{}, opt QueryOptions) (*gorm.DB, error) {
	result := tenancy.Db.Select(opt.Select).Where(opt.Where).Offset(*opt.Offset).Limit(*opt.Limit).Order(opt.Order).Find(&model)
	if result.Error != nil {
		return nil, result.Error
	}
	return res
}

func (tenancy *Tenancy) FindOne(model interface{}) *gorm.DB {

}
