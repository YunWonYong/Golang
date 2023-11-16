package core_router

import (
	"fmt"
	todo_router "gt/chap31/ex/todo/router"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Load(e *echo.Echo) {
	e.Pre(pre())
	todoGroup := e.Group("/todos")
	todo_router.Load(todoGroup)
}

func pre() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			header := c.Request().Header
			if header.Get("Content-Type") != "application/json" {
				header.Set("Content-Type", "application/json")
				c.Request().Header = header
			}

			err := next(c)

			if err != nil {
				fmt.Println(err.Error())
				c.JSON(http.StatusInternalServerError, map[string]interface{}{
					"err": err.Error(),
				})
			}

			header = c.Response().Header()
			if header.Get("Content-Type") != "application/json" {
				c.Response().Header().Set("Content-Type", "application/json")
			}
			return nil
		}
	}
}
