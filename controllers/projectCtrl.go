package controllers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/vitorsalgado/la-democracia/projects-reader/services"
)

type ProjectCtrl struct{}

func NewProjectCtrl() *ProjectCtrl {
	return &ProjectCtrl{}
}

func (ctrl *ProjectCtrl) List(c echo.Context) error {
	req := new(services.ProjectsRequest)

	if err := c.Bind(req); err != nil {
		return err
	}

	svc := services.NewProjectService()
	res := svc.List(req)

	if res.Success {
		return c.JSON(http.StatusOK, res)
	}

	return c.JSON(http.StatusBadRequest, res)
}
