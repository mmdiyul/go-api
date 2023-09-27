package common

import (
	"gorm.io/plugin/soft_delete"
	"time"
)

type BaseModel struct {
	ID        uint                  `json:"id" gorm:"primary_key"`
	CreatedAt time.Time             `json:"createdAt" gorm:"default:current_timestamp;not null"`
	UpdatedAt time.Time             `json:"updatedAt" gorm:"default:current_timestamp;not null"`
	CreatedBy uint                  `json:"createdBy"`
	UpdatedBy uint                  `json:"updatedBy"`
	DeletedAt time.Time             `json:"deletedAt"`
	Deleted   soft_delete.DeletedAt `json:"deleted" gorm:"softDelete:flag,DeletedAtField:DeletedAt"`
}
