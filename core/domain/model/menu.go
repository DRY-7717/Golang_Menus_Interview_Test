package model

import (
	"time"

	"github.com/google/uuid"
)

type Menu struct {
	ID        uuid.UUID  `gorm:"type:uuid;column:id;primaryKey;default:gen_random_uuid()"`
	MenuID    *uuid.UUID `gorm:"type:uuid;index;column:menu_id"`
	Name      string     `gorm:"column:name;not null"`
	Depth     int        `gorm:"column:depth"`
	SortOrder int        `gorm:"column:sort_order"`
	CreatedAt time.Time  `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time  `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

func (Menu) TableName() string {
	return "menus"
}
