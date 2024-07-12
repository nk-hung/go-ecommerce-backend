package po

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string    `gorm:"column:user_name"`
	Roles    []Role    `gorm:"many2many:go_user_role"`
	IsActive bool      `gorm:"column:is_active; type:boolean"`
	UUID     uuid.UUID `gorm:"column:uuid; type:varchar(255); not null; index idx_uuid; unique"`
}

func (u *User) TableName() string {
	return "go_db_user"
}
