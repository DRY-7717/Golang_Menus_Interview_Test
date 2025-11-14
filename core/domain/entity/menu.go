package entity

import (
	"github.com/google/uuid"
)

type MenuEntity struct {
	ID        uuid.UUID    `json:"id"`
	MenuID    *uuid.UUID   `json:"menu_id"`
	Name      string       `json:"name"`
	Depth     int          `json:"depth"`
	SortOrder int          `json:"sort_order"`
	Children  []MenuEntity `json:"children"`
}
