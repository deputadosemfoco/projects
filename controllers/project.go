package controllers

import (
	"net/http"

	"github.com/deputadosemfoco/projects/models"
	"github.com/labstack/echo"
)

type (
	ProjectRepository interface {
		List(ids []int64) []models.Project
	}

	ProjectCtrl struct {
		ProjectRepository ProjectRepository
	}
)

func (ctrl *ProjectCtrl) List(c echo.Context) error {
	req := new(models.ProjectsRequest)

	if err := c.Bind(req); err != nil {
		return err
	}

	res := ctrl.ProjectRepository.List(req.Projects)

	if len(res) > 0 {
		return c.JSON(http.StatusOK, res)
	}

	return c.JSON(http.StatusNoContent, nil)
}
