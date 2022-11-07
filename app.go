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
	token := tokenResp.Credential.Token

	return &App{
		MatchesEndpoint:   matches.NewEndpoint(Domain, token),
		StandingsEndpoint: standings.NewEndpoint(Domain, token),
		TeamEndpoint:      teams.NewEndpoint(Domain, token),
	}
}

func (a *App) Run() {
}
