package models

type (
	ProjectsRequest struct {
		Projects []int64
	}

	Project struct {
		ID              string `json:"id"`
		Name            string `json:"name"`
		Type            string `json:"type"`
		TypeDescription string `json:"typeDescription" db:"typeDescription"`
		Year            string `json:"year"`
		PresentedAt     string `json:"date" db:"presentedAt"`
		Briefing        string `json:"briefing"`
		Explanation     string `json:"explanation"`
		AuthorID        string `json:"authorId" db:"authorId"`
		Author          string `json:"author"`
		Party           string `json:"party"`
		PartyName       string `json:"partyName" db:"partyName"`
		Situation       string `json:"situation"`
	}
)
