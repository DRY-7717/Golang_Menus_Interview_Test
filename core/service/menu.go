package service

import (
	"context"
	"errors"
	"golang_menu_interview/core/domain/entity"
	"golang_menu_interview/internal/adapter/repository"
	"golang_menu_interview/utils/treemenu"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

type MenuServiceInterface interface {
	CreateMenu(ctx context.Context, req entity.MenuEntity) error
	FindAllMenu(ctx context.Context) ([]entity.MenuEntity, error)
	FindMenuByID(ctx context.Context, id uuid.UUID) (*entity.MenuEntity, error)
	UpdateMenu(ctx context.Context, req entity.MenuEntity) error
	DeleteMenu(ctx context.Context, id uuid.UUID) error
	MoveMenu(ctx context.Context, req entity.MenuEntity) error
	ReorderMenu(ctx context.Context, req entity.MenuEntity) error
}

type MenuService struct {
	MenuRepoInterface repository.MenuRepositoryInterface
}

func NewMenuService(menuRepoInterface repository.MenuRepositoryInterface) MenuServiceInterface {
	return &MenuService{
		MenuRepoInterface: menuRepoInterface,
	}
}

// CreateMenu implements MenuServiceInterface.
func (m *MenuService) CreateMenu(ctx context.Context, req entity.MenuEntity) error {

	if req.MenuID != nil {
		parent, err := m.MenuRepoInterface.FindMenuByID(ctx, *req.MenuID)
		if err != nil {
			log.Err(err).Msg("[SERVICE] CreateMenu - 1 ")
			return err
		}
		req.Depth = parent.Depth + 1
	} else {
		req.Depth = 0
	}

	return m.MenuRepoInterface.CreateMenu(ctx, req)
}

// FindAllMenu implements MenuServiceInterface.
func (m *MenuService) FindAllMenu(ctx context.Context) ([]entity.MenuEntity, error) {
	menus, err := m.MenuRepoInterface.FindAllMenu(ctx)
	if err != nil {
		log.Err(err).Msg("[SERVICE] GetAllMenus - 1")
		return nil, err
	}

	tree := treemenu.BuildTree(menus, nil)
	return tree, nil
}

// FindMenuByID implements MenuServiceInterface.
func (m *MenuService) FindMenuByID(ctx context.Context, id uuid.UUID) (*entity.MenuEntity, error) {
	menu, err := m.MenuRepoInterface.FindMenuByID(ctx, id)
	if err != nil {
		log.Err(err).Msg("[SERVICE] FindMenuByID - 1")
		return nil, err
	}

	allMenus, err := m.MenuRepoInterface.FindAllMenu(ctx)
	if err != nil {
		log.Err(err).Msg("[SERVICE] FindMenuByID - 2")
		return nil, err
	}

	menu.Children = treemenu.BuildTree(allMenus, &menu.ID)

	return menu, nil
}

// UpdateMenu implements MenuServiceInterface.
func (m *MenuService) UpdateMenu(ctx context.Context, req entity.MenuEntity) error {
	return m.MenuRepoInterface.UpdateMenu(ctx, req)
}

// DeleteMenu implements MenuServiceInterface.
func (m *MenuService) DeleteMenu(ctx context.Context, id uuid.UUID) error {
	return m.MenuRepoInterface.DeleteMenu(ctx, id)
}

// MoveMenu implements MenuServiceInterface.
func (m *MenuService) MoveMenu(ctx context.Context, req entity.MenuEntity) error {

	currentMenu, err := m.MenuRepoInterface.FindMenuByID(ctx, req.ID)
	if err != nil {
		log.Err(err).Msg("[SERVICE] MoveMenu - 1")
		return err
	}

	var newDepth int
	if req.MenuID != nil {

		parent, err := m.MenuRepoInterface.FindMenuByID(ctx, *req.MenuID)
		if err != nil {
			log.Err(err).Msg("[SERVICE] MoveMenu - 2")
			return err
		}

		isDesc, err := m.MenuRepoInterface.IsDescendant(ctx, parent.ID, req.ID)
		if err != nil {
			log.Err(err).Msg("[SERVICE] MoveMenu - 3")
			return err
		}
		if isDesc {
			return errors.New("cannot move menu to its own descendant")
		}

		newDepth = parent.Depth + 1
	} else {

		newDepth = 0
	}

	depthDiff := newDepth - currentMenu.Depth

	req.Depth = newDepth
	if err := m.MenuRepoInterface.MoveMenu(ctx, req); err != nil {
		log.Err(err).Msg("[SERVICE] MoveMenu - 4")
		return err
	}

	if depthDiff != 0 {
		if err := m.MenuRepoInterface.UpdateDescendantsDepth(ctx, req.ID, depthDiff); err != nil {
			log.Err(err).Msg("[SERVICE] MoveMenu - 5")
			return err
		}
	}

	return nil
}

func (m *MenuService) ReorderMenu(ctx context.Context, req entity.MenuEntity) error {

	_, err := m.MenuRepoInterface.FindMenuByID(ctx, req.ID)
	if err != nil {
		log.Err(err).Msg("[SERVICE] ReorderMenu - 1")
		return err
	}

	return m.MenuRepoInterface.ReorderMenu(ctx, req)
}
