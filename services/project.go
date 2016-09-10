package services

import "github.com/vitorsalgado/la-democracia/lib/go/messages"

type (
	ProjectService interface {
		List(req *ProjectsRequest) messages.Response
	}

	ProjectServiceImpl struct {
	}
)

func NewProjectService() *ProjectServiceImpl {
	return &ProjectServiceImpl{}
}

func (svc *ProjectServiceImpl) List(request *ProjectsRequest) messages.Response {
	reader := NewProjectReader()
	data := reader.List(request.Projects)

	return messages.Response{Success: true, Data: data}
}
