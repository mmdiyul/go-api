package models

import (
	"go-api/common"
)

type User struct {
	common.BaseModel
	FirstName string `json:"firstName" gorm:"size:100;not null"`
	LastName  string `json:"lastName" gorm:"size:100;not null"`
	Username  string `json:"username" gorm:"size:50;not null"`
	Email     string `json:"email" gorm:"size:100;not null"`
	Password  string `json:"password"`
	RoleId    uint   `json:"roleId" gorm:"not null"`
	Role      Role   `json:"role,omitempty" gorm:"foreignKey:RoleId;references:ID"`
}

func (User) TableName() string {
	return "sys_user"
}
