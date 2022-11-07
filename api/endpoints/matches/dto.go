package matches

type Match struct {
	ID           string `json:"id"`
	Type         string `json:"type"`
	Group        string `json:"group"`
	HomeTeamID   string `json:"home_team_id"`
	AwayTeamID   string `json:"away_team_id"`
	HomeScore    uint8  `json:"home_score"`
	AwayScore    uint8  `json:"away_score"`
	LocalDate    string `json:"local_date"`
	TimeElapsed  string `json:"time_elapsed"`
	Finished     string `json:"finished"`
	MatchDay     string `json:"matchday"`
	HomeTeamName string `json:"home_team_en"`
	AwayTeamName string `json:"away_team_en"`
	// HomeScorers  string  `json:"home_scorers"`
	// AwayScorers  string  `json:"away_scorers"`
}

type MatchesResponse struct {
	Status  string  `json:"status"`
	Matches []Match `json:"data"`
}
