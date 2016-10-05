package routes

import (
	"net/http"
	"os"
	"runtime"

	"bitbucket.org/astromatch/gemeos/messages"
	"github.com/deputadosemfoco/projects/controllers"
	"github.com/deputadosemfoco/projects/repositories"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func SetUp() *echo.Echo {
	e := echo.New()

	e.SetDebug(os.Getenv("GO_ENV") == "development")

	e.Use(middleware.Recover())
	e.SetHTTPErrorHandler(errorHandler)
	e.Use(middleware.Gzip())

	healthCtrl := controllers.HealthCtrl{}

	e.Get("/api/chk", healthCtrl.Check)

	projectCtrl := controllers.ProjectCtrl{}
	projectCtrl.ProjectRepository = new(repositories.ProjectRepository)

	e.Post("/api/project/list", projectCtrl.List)

	return e
}

func errorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	msg := err.Error()
	detail := ""

	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		msg = he.Message
	}

	if msg == "" {
		msg = http.StatusText(code)
	}

	if os.Getenv("GO_ENV") == "development" {
		b := make([]byte, 2048)
		n := runtime.Stack(b, false)

		detail = string(b[:n])
	}

	resp := messages.Error{Message: msg, Code: code, Detail: detail}

	if !c.Response().Committed() {
		c.JSON(code, resp)
	}
}
