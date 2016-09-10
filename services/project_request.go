package services

import "github.com/vitorsalgado/la-democracia/lib/go/messages"

type ProjectsRequest struct {
	messages.Request

	Projects []int64
}
