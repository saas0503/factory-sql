package tenant

import "github.com/saas0503/fsql/base"

type Tenant struct {
	base.Model `gorm:"embedded"`
	Name       string `gorm:"varchar(100);not null"`
	Code       string `gorm:"varchar(100);not null;uniqueIndex:idx_code"`
}
