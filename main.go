package main

import (
	"fmt"
	"github.com/guidoenr/fulbo/model"
	"github.com/guidoenr/fulbo/model/psdb"
)

func main() {
	var match model.Match

	err := match.InitFromJson()
	if err != nil {
		fmt.Printf("error: %v", err)
	}

	match.GenerateTeams()

	match.Team1.Show()
	match.Team2.Show()
	match.GenerateTeams()
	var db psdb.PostgreDB
	db.InitDB()
}
