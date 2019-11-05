package main

import (
	"os"

	"github.com/davepmiller/todomvc/controller"
	"github.com/davepmiller/todomvc/database"
	"github.com/labstack/echo"
	_ "github.com/lib/pq"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	db := database.Init()
	e := echo.New()
	e.GET("/todo", controller.GetAll(db))
	e.PUT("/todo", controller.Put(db))
	e.DELETE("/todo/:id", controller.Delete(db))
	e.Logger.Fatal(e.Start(":" + port))
}
