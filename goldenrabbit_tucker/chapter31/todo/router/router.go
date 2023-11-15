package todo_router

import (
	todo_handler "gt/chap31/ex/todo/handler"

	"github.com/labstack/echo/v4"
)

func Load(e *echo.Group) {
	e.GET("/:writer", todo_handler.FindAllHandler)
	e.POST("", todo_handler.RegistHandler)
	e.PUT("/:writer/:id", todo_handler.ModifyHandler)
	e.PATCH("/:writer/:id", todo_handler.ClearFlagToggleHandler)
	e.DELETE("/:writer/:id", todo_handler.DestroyHandler)
}
