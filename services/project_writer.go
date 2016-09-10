package services

import (
	"fmt"
	"log"

	"github.com/vitorsalgado/la-democracia/lib/go/caching/keys"
	"github.com/vitorsalgado/la-democracia/lib/go/redisdb"
	"github.com/vitorsalgado/la-democracia/projects-reader/models"
)

func WriteProject(projects []models.Project) {
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
