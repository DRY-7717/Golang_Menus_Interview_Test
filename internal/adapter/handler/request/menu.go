package request

type MenuRequest struct {
	MenuID    string `json:"menu_id"`
	Name      string `json:"name"  validate:"required"`
	SortOrder int    `json:"sort_order"`
}

type MoveMenuRequest struct {
	NewMenuID string `json:"new_menu_id"`
}

type ReorderMenuRequest struct {
	NewSortOrder int `json:"new_sort_order"`
}
