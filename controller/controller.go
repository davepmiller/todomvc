package controller

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/davepmiller/todomvc/model"
	"github.com/davepmiller/todomvc/types"
	"github.com/labstack/echo"
)

func GetAll(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, model.GetAll(db))
	}
}

func Put(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var todo types.Todo
		c.Bind(&todo)
		id, err := model.Put(db, todo.Description)
		if err == nil {
			return c.JSON(
				http.StatusCreated,
				types.Response{"created": id})
		} else {
			return err
		}
	}
}

func Delete(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		_, err := model.Delete(db, id)
		if err == nil {
			return c.JSON(
				http.StatusOK,
				types.Response{"deleted": id})
		} else {
			return err
		}
	}
}
