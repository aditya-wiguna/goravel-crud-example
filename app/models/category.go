package models

import (
	"github.com/goravel/framework/database/orm"
)

type Category struct {
	orm.Model
	Name string
}
