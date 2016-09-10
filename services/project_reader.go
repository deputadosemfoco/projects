package services

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/vitorsalgado/la-democracia/lib/go/caching/keys"
	"github.com/vitorsalgado/la-democracia/lib/go/redisdb"
	"github.com/vitorsalgado/la-democracia/lib/go/sqldb"
	"github.com/vitorsalgado/la-democracia/projects-reader/models"
)

type (
	ProjectReader interface {
		List(ids []int) []models.Project
	}

	ProjectReaderImpl struct {
	}
)

func NewProjectReader() *ProjectReaderImpl {
	return &ProjectReaderImpl{}
}

func (projectReader *ProjectReaderImpl) List(ids []int64) []models.Project {
	list := fetchFromRedis(ids)

	if list == nil || len(list) == 0 {
		list = fetchFromSQL(ids)

		if list != nil && len(list) > 0 {
			WriteProject(list)
		}
	}

	return list
}

func fetchFromRedis(ids []int64) []models.Project {
	redis, _ := redisdb.Connect()
	data := []models.Project{}

	for _, key := range ids {
		redis.PipeAppend("HGETALL", fmt.Sprintf("%s:%s", keys.ProjectKEY, key))
	}

	for {
		resp, err := redis.PipeResp().Map()

		if err != nil {
			break
		}

		if resp[keys.ProjectID] == "" {
			break
		}

		project := models.Project{
			ID:              resp[keys.ProjectID],
			Name:            resp[keys.ProjectName],
			Type:            resp[keys.ProjectType],
			TypeDescription: resp[keys.ProjectTypeDescription],
			Year:            resp[keys.ProjectYear],
			PresentedAt:     resp[keys.ProjectPresentedAt],
			Briefing:        resp[keys.ProjectBriefing],
			Explanation:     resp[keys.ProjectExplanation],
			AuthorID:        resp[keys.ProjectAuthorID],
			Author:          resp[keys.ProjectAuthor],
			Party:           resp[keys.ProjectParty],
			PartyName:       resp[keys.ProjectPartyName],
			Situation:       resp[keys.ProjectSituation],
		}

		data = append(data, project)
	}

	return data
}

func fetchFromSQL(ids []int64) []models.Project {
	db := sqldb.Connect()
	resp := []models.Project{}

	sql := fmt.Sprintf(
		"SELECT id, name, `type`, typeDescription, `year`, presentedAt, briefing, explanation, authorId, author author, party, partyName, situation FROM flat_project WHERE id IN (?);")

	query, args, _ := sqlx.In(sql, ids)
	query = db.Rebind(query)
	db.Select(&resp, query, args...)

	return resp
}
