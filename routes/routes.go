package routes

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/vitorsalgado/la-democracia/projects-reader/controllers"
)

type (
	Router struct{}
)

func (router *Router) SetUp() *echo.Echo {
	e := echo.New()
	e.Use(middleware.Gzip())

	healthCtrl := controllers.NewHealthCtrl()
	projectCtrl := controllers.NewProjectCtrl()

	e.SetHTTPErrorHandler(controllers.GtHTTPErrorHandler)

	e.Get("/api/chk", healthCtrl.Check)

	e.Post("/api/project/list", projectCtrl.List)

	return e
}
