package handler

import (
	"golang_menu_interview/core/domain/entity"
	"golang_menu_interview/core/service"
	"golang_menu_interview/internal/adapter/handler/request"
	"golang_menu_interview/internal/adapter/handler/response"
	"golang_menu_interview/utils/validation"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

type MenuHandlerInterface interface {
	CreateMenu(c *fiber.Ctx) error
	FindAllMenu(c *fiber.Ctx) error
	FindMenuByID(c *fiber.Ctx) error
	UpdateMenu(c *fiber.Ctx) error
	DeleteMenu(c *fiber.Ctx) error
	MoveMenu(c *fiber.Ctx) error
	ReorderMenu(c *fiber.Ctx) error
}

type MenuHandler struct {
	MenuServiceInterface service.MenuServiceInterface
	Validator            *validator.Validate
}

func NewMenuHandler(menuServiceInterface service.MenuServiceInterface, validator *validator.Validate) MenuHandlerInterface {
	return &MenuHandler{
		MenuServiceInterface: menuServiceInterface,
		Validator:            validator,
	}
}

// CreateMenu implements MenuHandlerInterface.
func (m *MenuHandler) CreateMenu(c *fiber.Ctx) error {
	var (
		req     = request.MenuRequest{}
		resp    = response.SuccessResponseDefault{}
		respErr = response.ErrorResponseDefault{}
		ctx     = c.UserContext()
	)

	if err := c.BodyParser(&req); err != nil {
		log.Error().Err(err).Msg("[HANDLER] CreateMenu - 1 ")
		respErr.Message = err.Error()
		respErr.Status = false
		return c.Status(fiber.StatusUnprocessableEntity).JSON(respErr)
	}

	if err := m.Validator.Struct(&req); err != nil {
		log.Error().Err(err).Msg("[HANDLER] CreateMenu - 2 ")
		errors := validation.CustomValidator(err)
		respErr.Message = "Invalid request"
		respErr.Status = false
		respErr.Errors = errors
		return c.Status(fiber.StatusBadRequest).JSON(respErr)
	}

	var reqEntity entity.MenuEntity

	if req.MenuID != "" {
		menuUUID, err := uuid.Parse(req.MenuID)
		if err != nil {
			log.Error().Err(err).Msg("[HANDLER] CreateMenu - invalid menu_id")
			respErr.Message = "Invalid menu_id format"
			respErr.Status = false
			return c.Status(fiber.StatusBadRequest).JSON(respErr)
		}
		reqEntity.MenuID = &menuUUID
	}

	reqEntity.Name = req.Name
	reqEntity.SortOrder = req.SortOrder

	if err := m.MenuServiceInterface.CreateMenu(ctx, reqEntity); err != nil {
		log.Error().Err(err).Msg("[HANDLER] CreateCategory - 3")
		respErr.Message = err.Error()
		respErr.Status = false
		return c.Status(fiber.StatusInternalServerError).JSON(respErr)
	}

	resp.Message = "Create menu successfully"
	resp.Status = true
	resp.Data = nil
	return c.Status(fiber.StatusCreated).JSON(resp)

}

// FindAllMenu implements MenuHandlerInterface.
func (m *MenuHandler) FindAllMenu(c *fiber.Ctx) error {
	var (
		resp    = response.SuccessResponseDefault{}
		respErr = response.ErrorResponseDefault{}
		ctx     = c.UserContext()
	)

	menus, err := m.MenuServiceInterface.FindAllMenu(ctx)
	if err != nil {
		log.Error().Err(err).Msg("[HANDLER] FindAllMenu - 1")
		respErr.Message = err.Error()
		respErr.Status = false
		return c.Status(fiber.StatusInternalServerError).JSON(respErr)
	}

	resp.Message = "Find all menus successfully"
	resp.Status = true
	resp.Data = menus
	return c.Status(fiber.StatusOK).JSON(resp)
}

func (m *MenuHandler) FindMenuByID(c *fiber.Ctx) error {
	var (
		resp    = response.SuccessResponseDefault{}
		respErr = response.ErrorResponseDefault{}
		ctx     = c.UserContext()
	)

	idMenu := c.Params("id")
	id, err := uuid.Parse(idMenu)

	if err != nil {
		log.Error().Err(err).Msg("[HANDLER] FindMenuByID - 1")
		respErr.Message = "Invalid menu ID format"
		respErr.Status = false
		return c.Status(fiber.StatusBadRequest).JSON(respErr)
	}

	menu, err := m.MenuServiceInterface.FindMenuByID(ctx, id)
	if err != nil {
		log.Error().Err(err).Msg("[HANDLER] FindMenuByID - 2")
		respErr.Message = err.Error()
		respErr.Status = false
		return c.Status(fiber.StatusInternalServerError).JSON(respErr)
	}

	resp.Message = "Find menu by id successfully"
	resp.Status = true
	resp.Data = menu
	return c.Status(fiber.StatusOK).JSON(resp)
}

// UpdateMenu implements MenuHandlerInterface.
func (m *MenuHandler) UpdateMenu(c *fiber.Ctx) error {

	var (
		req     = request.MenuRequest{}
		resp    = response.SuccessResponseDefault{}
		respErr = response.ErrorResponseDefault{}
		ctx     = c.UserContext()
	)

	if err := c.BodyParser(&req); err != nil {
		log.Error().Err(err).Msg("[HANDLER] UpdateMenu - 1 ")
		respErr.Message = err.Error()
		respErr.Status = false
		return c.Status(fiber.StatusUnprocessableEntity).JSON(respErr)
	}

	if err := m.Validator.Struct(&req); err != nil {
		log.Error().Err(err).Msg("[HANDLER] UpdateMenu - 2 ")
		errors := validation.CustomValidator(err)
		respErr.Message = "Invalid request"
		respErr.Status = false
		respErr.Errors = errors
		return c.Status(fiber.StatusBadRequest).JSON(respErr)
	}

	idMenu := c.Params("id")
	id, err := uuid.Parse(idMenu)

	if err != nil {
		log.Error().Err(err).Msg("[HANDLER] UpdateMenu - 3")
		respErr.Message = "Invalid menu ID format"
		respErr.Status = false
		return c.Status(fiber.StatusBadRequest).JSON(respErr)
	}

	var reqEntity entity.MenuEntity

	reqEntity.ID = id
	reqEntity.Name = req.Name
	reqEntity.SortOrder = req.SortOrder

	if err := m.MenuServiceInterface.UpdateMenu(ctx, reqEntity); err != nil {
		log.Error().Err(err).Msg("[HANDLER] UpdateMenu - 4")
		respErr.Message = err.Error()
		respErr.Status = false
		return c.Status(fiber.StatusInternalServerError).JSON(respErr)
	}

	resp.Message = "Update menu successfully"
	resp.Status = true
	resp.Data = nil
	return c.Status(fiber.StatusOK).JSON(resp)

}

// DeleteMenu implements MenuHandlerInterface.
func (m *MenuHandler) DeleteMenu(c *fiber.Ctx) error {
	var (
		resp    = response.SuccessResponseDefault{}
		respErr = response.ErrorResponseDefault{}
		ctx     = c.UserContext()
	)

	idMenu := c.Params("id")
	id, err := uuid.Parse(idMenu)

	if err != nil {
		log.Error().Err(err).Msg("[HANDLER] DeleteMenu - 1")
		respErr.Message = "Invalid menu ID format"
		respErr.Status = false
		return c.Status(fiber.StatusBadRequest).JSON(respErr)
	}

	if err := m.MenuServiceInterface.DeleteMenu(ctx, id); err != nil {
		log.Error().Err(err).Msg("[HANDLER] DeleteMenu - 2")
		respErr.Message = err.Error()
		respErr.Status = false
		return c.Status(fiber.StatusNotFound).JSON(respErr)
	}

	resp.Message = "Delete menu successfully"
	resp.Status = true
	resp.Data = nil
	return c.Status(fiber.StatusOK).JSON(resp)
}

// MoveMenu implements MenuHandlerInterface.
func (m *MenuHandler) MoveMenu(c *fiber.Ctx) error {
	var (
		req     = request.MoveMenuRequest{}
		resp    = response.SuccessResponseDefault{}
		respErr = response.ErrorResponseDefault{}
		ctx     = c.UserContext()
	)

	idMenu := c.Params("id")
	id, err := uuid.Parse(idMenu)
	if err != nil {
		log.Error().Err(err).Msg("[HANDLER] MoveMenu - 1")
		respErr.Message = "Invalid menu ID format"
		respErr.Status = false
		return c.Status(fiber.StatusBadRequest).JSON(respErr)
	}

	if err := c.BodyParser(&req); err != nil {
		log.Error().Err(err).Msg("[HANDLER] MoveMenu - 2")
		respErr.Message = err.Error()
		respErr.Status = false
		return c.Status(fiber.StatusUnprocessableEntity).JSON(respErr)
	}

	var reqEntity entity.MenuEntity
	reqEntity.ID = id

	if req.NewMenuID != "" {
		menuUUID, err := uuid.Parse(req.NewMenuID)
		if err != nil {
			log.Error().Err(err).Msg("[HANDLER] MoveMenu - 3")
			respErr.Message = "Invalid new_menu_id format"
			respErr.Status = false
			return c.Status(fiber.StatusBadRequest).JSON(respErr)
		}
		reqEntity.MenuID = &menuUUID
	} else {
		reqEntity.MenuID = nil
	}

	if err := m.MenuServiceInterface.MoveMenu(ctx, reqEntity); err != nil {
		log.Error().Err(err).Msg("[HANDLER] MoveMenu - 4")

		status := fiber.StatusInternalServerError
		if err.Error() == "menu not found" {
			status = fiber.StatusNotFound
		} else if err.Error() == "cannot move menu to its own descendant" {
			status = fiber.StatusBadRequest
		}

		respErr.Message = err.Error()
		respErr.Status = false
		return c.Status(status).JSON(respErr)
	}

	resp.Message = "Move menu successfully"
	resp.Status = true
	resp.Data = nil
	return c.Status(fiber.StatusOK).JSON(resp)
}

// ReorderMenu implements MenuHandlerInterface.
func (m *MenuHandler) ReorderMenu(c *fiber.Ctx) error {
	var (
		req     = request.ReorderMenuRequest{}
		resp    = response.SuccessResponseDefault{}
		respErr = response.ErrorResponseDefault{}
		ctx     = c.UserContext()
	)

	idMenu := c.Params("id")
	id, err := uuid.Parse(idMenu)
	if err != nil {
		log.Error().Err(err).Msg("[HANDLER] ReorderMenu - 1")
		respErr.Message = "Invalid menu ID format"
		respErr.Status = false
		return c.Status(fiber.StatusBadRequest).JSON(respErr)
	}

	if err := c.BodyParser(&req); err != nil {
		log.Error().Err(err).Msg("[HANDLER] ReorderMenu - 2")
		respErr.Message = err.Error()
		respErr.Status = false
		return c.Status(fiber.StatusUnprocessableEntity).JSON(respErr)
	}

	if err := m.Validator.Struct(&req); err != nil {
		log.Error().Err(err).Msg("[HANDLER] ReorderMenu - 3")
		errors := validation.CustomValidator(err)
		respErr.Message = "Invalid request"
		respErr.Status = false
		respErr.Errors = errors
		return c.Status(fiber.StatusBadRequest).JSON(respErr)
	}

	var reqEntity entity.MenuEntity
	reqEntity.ID = id
	reqEntity.SortOrder = req.NewSortOrder

	if err := m.MenuServiceInterface.ReorderMenu(ctx, reqEntity); err != nil {
		log.Error().Err(err).Msg("[HANDLER] ReorderMenu - 4")

		status := fiber.StatusInternalServerError
		if err.Error() == "menu not found" {
			status = fiber.StatusNotFound
		}

		respErr.Message = err.Error()
		respErr.Status = false
		return c.Status(status).JSON(respErr)
	}

	resp.Message = "Reorder menu successfully"
	resp.Status = true
	resp.Data = nil
	return c.Status(fiber.StatusOK).JSON(resp)
}
