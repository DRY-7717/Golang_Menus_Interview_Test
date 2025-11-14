package router

import (
	"golang_menu_interview/core/service"
	"golang_menu_interview/internal/adapter/handler"
	"golang_menu_interview/internal/adapter/repository"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func MenuRouter(api fiber.Router, db *gorm.DB, validator *validator.Validate) {

	menuRepository := repository.NewMenuRepository(db)
	menuService := service.NewMenuService(menuRepository)
	menuHandler := handler.NewMenuHandler(menuService, validator)

	api.Get("/menus", menuHandler.FindAllMenu)
	api.Get("/menus/:id", menuHandler.FindMenuByID)
	api.Post("/menus", menuHandler.CreateMenu)
	api.Put("/menus/:id", menuHandler.UpdateMenu)
	api.Delete("/menus/:id", menuHandler.DeleteMenu)
	api.Patch("/menus/:id/move", menuHandler.MoveMenu)
	api.Patch("/menus/:id/reorder", menuHandler.ReorderMenu)
}
