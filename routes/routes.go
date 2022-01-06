package routes

import (
	"crud-movies/config"
	"crud-movies/controllers"

	"github.com/labstack/echo/v4"
)

func InitRoute() *echo.Echo {
	//init echo
	e := echo.New()

	//Database
	db := config.CreateConnection()

	//movie routes
	e.GET("/articles", controllers.GetArticless(db))
	e.POST("/articles", controllers.AddArticles(db))
	e.PUT("/articles/:id", controllers.UpdateArticles(db))
	e.DELETE("/articles/:id", controllers.DeleteArticles(db))

	return e
}
