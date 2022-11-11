package main

import (
	"os"

	"github.com/adalessa/qatar/api/endpoints/auth"
	"github.com/adalessa/qatar/api/endpoints/matches"
	"github.com/adalessa/qatar/api/endpoints/standings"
	"github.com/adalessa/qatar/api/endpoints/teams"
)

const Domain = "http://api.cup2022.ir"

type App struct {
	MatchesEndpoint   *matches.MatchesEndpoint
	StandingsEndpoint *standings.StandingsEndpoint
	TeamEndpoint      *teams.TeamsEndpoint
}

func NewApp() *App {
	var token string
	tokenRaw, err := os.ReadFile("/tmp/qatar_token")
	if err != nil {
		authEndpoint := auth.NewEndpoint(Domain)

		email := os.Getenv("QATAR_EMAIL")
		password := os.Getenv("QATAR_PASSWORD")
		if email == "" || password == "" {
			panic("no credentials")
		}

		tokenResp, err := authEndpoint.Login(email, password)
		if err != nil {
			panic(err)
		}
		if tokenResp.Status != "success" {
			panic("cant get the token")
		}
		token = tokenResp.Credential.Token
		err = os.WriteFile("/tmp/qatar_token", []byte(token), 0644)
		if err != nil {
			panic(err)
		}
	} else {
		token = string(tokenRaw)
	}

	return &App{
		MatchesEndpoint:   matches.NewEndpoint(Domain, token),
		StandingsEndpoint: standings.NewEndpoint(Domain, token),
		TeamEndpoint:      teams.NewEndpoint(Domain, token),
	}
}

func (a *App) Run() {
	resp, err := a.MatchesEndpoint.GetByMatchDayID("3")
	if err != nil {
		panic(err)
	}

	println(resp.Matches)
}
