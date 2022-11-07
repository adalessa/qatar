package teams

type Team struct {
	ID       string `json:"id"`
	Name     string `json:"name_en"`
	Flag     string `json:"flag"`
	FifaCode string `json:"fifa_code"`
	Groups   string `json:"groups"`
}

type GetTeamsResponse struct {
	Status string `json:"status"`
	Data   []Team `json:"data"`
}
