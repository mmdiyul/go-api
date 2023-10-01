package models

import "go-api/common"

type Role struct {
	common.BaseModel
	RoleCode string `json:"roleCode" gorm:"size:4;not null"`
	RoleName string `json:"roleName" gorm:"size:20;not null"`
}

func (Role) TableName() string {
	return "sys_role"
}
