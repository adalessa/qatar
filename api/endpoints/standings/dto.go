package standings

type StandingResponse struct {
	Status    string     `json:"status"`
	Standings []Standing `json:"data"`
}

type Standing struct {
	Group        string         `json:"group"`
	TeamStanding []TeamStanding `json:"teams"`
}

type TeamStanding struct {
	TeamID          string `json:"team_id"`
	MatchesPlayed   string `json:"mp"`
	MatchesWon      string `json:"w"`
	MatchesLost     string `json:"l"`
	Score           string `json:"pts"`
	GoalsFor        string `json:"gf"`
	GoalsAgains     string `json:"ga"`
	GoalsDifference string `json:"gd"`
	Drawn           string `json:"d"`
	Name            string `json:"name_en"`
	Flag            string `json:"flag"`
}
