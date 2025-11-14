package repository

import (
	"context"
	"golang_menu_interview/core/domain/entity"
	"golang_menu_interview/core/domain/model"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type MenuRepositoryInterface interface {
	CreateMenu(ctx context.Context, req entity.MenuEntity) error
	FindAllMenu(ctx context.Context) ([]entity.MenuEntity, error)
	FindMenuByID(ctx context.Context, id uuid.UUID) (*entity.MenuEntity, error)
	UpdateMenu(ctx context.Context, req entity.MenuEntity) error
	DeleteMenu(ctx context.Context, id uuid.UUID) error
	MoveMenu(ctx context.Context, req entity.MenuEntity) error
	ReorderMenu(ctx context.Context, req entity.MenuEntity) error
	IsDescendant(ctx context.Context, targetID, menuID uuid.UUID) (bool, error)
	UpdateDescendantsDepth(ctx context.Context, menuID uuid.UUID, depthDiff int) error
}

type MenuRepository struct {
	DB *gorm.DB
}

func NewMenuRepository(db *gorm.DB) MenuRepositoryInterface {
	return &MenuRepository{
		DB: db,
	}
}

// CreateMenu implements MenuRepositoryInterface.
func (m *MenuRepository) CreateMenu(ctx context.Context, req entity.MenuEntity) error {

	modelMenu := model.Menu{
		MenuID:    req.MenuID,
		Name:      req.Name,
		Depth:     req.Depth,
		SortOrder: req.SortOrder,
	}

	if err := m.DB.Create(&modelMenu).Error; err != nil {
		log.Err(err).Msg("[REPOSITORY] CreateMenu - 1 ")
		return err
	}

	return nil

}

// FindAllMenu implements MenuRepositoryInterface.
func (m *MenuRepository) FindAllMenu(ctx context.Context) ([]entity.MenuEntity, error) {
	modelMenu := []model.Menu{}

	if err := m.DB.Select("id", "menu_id", "name", "depth", "sort_order").Find(&modelMenu).Order("sort_order ASC").Error; err != nil {
		log.Err(err).Msg("[REPOSITORY] FindAllMenu - 1")
		return nil, err
	}

	var menuEntities []entity.MenuEntity
	for _, data := range modelMenu {
		menuEntities = append(menuEntities, entity.MenuEntity{
			ID:        data.ID,
			MenuID:    data.MenuID,
			Name:      data.Name,
			Depth:     data.Depth,
			SortOrder: data.SortOrder,
		})
	}

	return menuEntities, nil
}

// FindHeroSectionByID implements MenuRepositoryInterface.
func (m *MenuRepository) FindMenuByID(ctx context.Context, id uuid.UUID) (*entity.MenuEntity, error) {
	modelMenu := model.Menu{}

	if err := m.DB.Select("id", "menu_id", "name", "depth", "sort_order").Where("id = ?", id).First(&modelMenu).Error; err != nil {
		log.Err(err).Msg("[REPOSITORY] FindMenuByID - 1 ")
		return nil, err
	}

	return &entity.MenuEntity{
		ID:        modelMenu.ID,
		MenuID:    modelMenu.MenuID,
		Name:      modelMenu.Name,
		Depth:     modelMenu.Depth,
		SortOrder: modelMenu.SortOrder,
	}, nil

}

// UpdateMenu implements MenuRepositoryInterface.
func (m *MenuRepository) UpdateMenu(ctx context.Context, req entity.MenuEntity) error {
	modelMenu := model.Menu{}

	if err := m.DB.Where("id = ?", req.ID).First(&modelMenu).Error; err != nil {
		log.Err(err).Msg("[REPOSITORY] UpdateMenu - 1 ")
		return err
	}

	modelMenu.Name = req.Name
	modelMenu.SortOrder = req.SortOrder

	if err := m.DB.Save(&modelMenu).Error; err != nil {
		log.Err(err).Msg("[REPOSITORY] UpdateMenu - 2 ")
		return err
	}

	return nil
}

// DeleteCategory implements MenuRepositoryInterface.
func (m *MenuRepository) DeleteMenu(ctx context.Context, id uuid.UUID) error {
	modelMenu := model.Menu{}

	if err := m.DB.Where("id = ?", id).First(&modelMenu).Error; err != nil {
		log.Err(err).Msg("[REPOSITORY] DeleteCategory - 1 ")
		return err
	}

	if err := m.DB.Delete(&modelMenu).Error; err != nil {
		log.Err(err).Msg("[REPOSITORY] DeleteCategory - 2 ")
		return err
	}

	return nil
}

// MoveMenu implements MenuRepositoryInterface.
func (m *MenuRepository) MoveMenu(ctx context.Context, req entity.MenuEntity) error {
	modelMenu := model.Menu{}

	if err := m.DB.Where("id = ?", req.ID).First(&modelMenu).Error; err != nil {
		log.Err(err).Msg("[REPOSITORY] MoveMenu - 1")
		return err
	}

	modelMenu.MenuID = req.MenuID
	modelMenu.Depth = req.Depth

	if err := m.DB.Save(&modelMenu).Error; err != nil {
		log.Err(err).Msg("[REPOSITORY] MoveMenu - 2")
		return err
	}

	return nil
}

// ReorderMenu implements MenuRepositoryInterface.
func (m *MenuRepository) ReorderMenu(ctx context.Context, req entity.MenuEntity) error {
	modelMenu := model.Menu{}

	if err := m.DB.Where("id = ?", req.ID).First(&modelMenu).Error; err != nil {
		log.Err(err).Msg("[REPOSITORY] ReorderMenu - 1")
		return err
	}

	modelMenu.SortOrder = req.SortOrder

	if err := m.DB.Save(&modelMenu).Error; err != nil {
		log.Err(err).Msg("[REPOSITORY] ReorderMenu - 2")
		return err
	}

	return nil
}

// IsDescendant implements MenuRepositoryInterface.
func (m *MenuRepository) IsDescendant(ctx context.Context, targetID, menuID uuid.UUID) (bool, error) {
	query := `
		WITH RECURSIVE descendants AS (
			SELECT id, menu_id FROM menus WHERE id = $1
			UNION ALL
			SELECT m.id, m.menu_id FROM menus m
			INNER JOIN descendants d ON m.menu_id = d.id
		)
		SELECT EXISTS(SELECT 1 FROM descendants WHERE id = $2)
	`

	var exists bool
	if err := m.DB.Raw(query, menuID, targetID).Scan(&exists).Error; err != nil {
		log.Err(err).Msg("[REPOSITORY] IsDescendant - 1")
		return false, err
	}

	return exists, nil
}

// UpdateDescendantsDepth implements MenuRepositoryInterface.
func (m *MenuRepository) UpdateDescendantsDepth(ctx context.Context, menuID uuid.UUID, depthDiff int) error {
	query := `
		WITH RECURSIVE descendants AS (
			SELECT id, depth FROM menus WHERE menu_id = $1
			UNION ALL
			SELECT m.id, m.depth FROM menus m
			INNER JOIN descendants d ON m.menu_id = d.id
		)
		UPDATE menus SET depth = depth + $2 WHERE id IN (SELECT id FROM descendants)
	`

	if err := m.DB.Exec(query, menuID, depthDiff).Error; err != nil {
		log.Err(err).Msg("[REPOSITORY] UpdateDescendantsDepth - 1")
		return err
	}

	return nil
}
