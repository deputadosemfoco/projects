package repositories

import (
	"fmt"
	"log"

	"github.com/deputadosemfoco/go-libs/caching/keys"
	"github.com/deputadosemfoco/go-libs/redisdb"
	"github.com/deputadosemfoco/projects/models"
	"github.com/jmoiron/sqlx"
)

type ProjectRepository struct{}

func (repository *ProjectRepository) List(ids []int64) []models.Project {
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
	db := Connect()
	resp := []models.Project{}

	sql := fmt.Sprintf(
		"SELECT id, name, `type`, typeDescription, `year`, presentedAt, briefing, explanation, authorId, author author, party, partyName, situation FROM flat_project WHERE id IN (?);")

	query, args, _ := sqlx.In(sql, ids)
	query = db.Rebind(query)
	db.Select(&resp, query, args...)

	return resp
}

func writeProject(projects []models.Project) {
	c, _ := redisdb.Connect()

	for _, project := range projects {
		c.PipeAppend("HMSET", fmt.Sprintf("%s:%s", keys.ProjectKEY, project.ID),
			keys.ProjectID, project.ID,
			keys.ProjectName, project.Name,
			keys.ProjectBriefing, project.Briefing,
			keys.ProjectExplanation, project.Explanation,
			keys.ProjectYear, project.Year,
			keys.ProjectType, project.Type,
			keys.ProjectTypeDescription, project.TypeDescription,
			keys.ProjectAuthorID, project.AuthorID,
			keys.ProjectAuthor, project.Author,
			keys.ProjectParty, project.Party,
			keys.ProjectPartyName, project.PartyName,
			keys.ProjectPresentedAt, project.PresentedAt,
			keys.ProjectSituation, project.Situation)
	}

	err := c.PipeResp().Err

	if err != nil {
		log.Println(err)
	}
}
