package po

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	ID       int64  `gorm:"column:uuid; type:varchar(255); not null; primaryKey; autoIncrement; comment:'primary key is id'`
	RoleName string `gorm:"column:role_name"`
	RoleNote string `gorm:"column:role_note; type:text;"`
}

func (u *Role) TableName() string {
	return "go_db_role"
}
