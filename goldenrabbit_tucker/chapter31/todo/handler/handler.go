package todo_handler

import (
	todo_dao "gt/chap31/ex/todo/dao"
	todo_model "gt/chap31/ex/todo/model"
	"gt/chap31/ex/util"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

func FindAllHandler(c echo.Context) error {
	writer := c.Param("writer")

	buff, err := todo_dao.List(writer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "todolist data not found.")
		return err
	}
	return c.JSON(http.StatusOK, string(buff))
}

func RegistHandler(c echo.Context) error {
	dto, buff, err := util.BindBuff[todo_model.TodoRegistDTO](c.Request().Body)
	if err != nil {
		c.String(http.StatusInternalServerError, "body to dto fail.")
		return err
	}

	if err := todo_dao.Insert(dto); err != nil {
		c.String(http.StatusInternalServerError, "regist fail.")
		return errors.WithMessagef(err, "todolist regist fail. requestBody: %s", buff)
	}

	return c.String(http.StatusOK, "ok")
}

func ClearFlagToggleHandler(c echo.Context) error {
	writer := c.Param("writer")
	id := c.Param("id")
	if !util.RequiredValueValidation(writer, id) {
		return errors.New("writer and id required")
	}

	if err := todo_dao.Toggle(writer, id); err != nil {
		return errors.WithMessage(err, "toggle fail.")
	}

	return c.String(http.StatusOK, "ok")
}

func ModifyHandler(c echo.Context) error {
	return nil
}

func DestroyHandler(c echo.Context) error {
	writer := c.Param("writer")
	id := c.Param("id")
	if !util.RequiredValueValidation(writer, id) {
		return errors.New("writer and id required")
	}

	if err := todo_dao.Delete(writer, id); err != nil {
		return errors.WithMessage(err, "delete fail.")
	}

	return c.String(http.StatusOK, "ok")
}
