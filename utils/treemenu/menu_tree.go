package treemenu

import (
	"golang_menu_interview/core/domain/entity"

	"github.com/google/uuid"
)

func BuildTree(allMenus []entity.MenuEntity, MenuID *uuid.UUID) []entity.MenuEntity {
	var result []entity.MenuEntity

	for _, menu := range allMenus {
		// Check if this menu belongs to current parent
		if (MenuID == nil && menu.MenuID == nil) ||
			(MenuID != nil && menu.MenuID != nil && *menu.MenuID == *MenuID) {

			// Recursively build children
			menu.Children = BuildTree(allMenus, &menu.ID)
			result = append(result, menu)
		}
	}

	return result
}
