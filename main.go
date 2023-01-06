package main

import "github.com/guidoenr/fulbo/model/psdb"

func main() {
	//var match model.Match
	//
	//err := match.InitFromJson()
	//if err != nil {
	//	fmt.Printf("error: %v", err)
	//}
	//
	//match.GenerateTeams()
	//
	//fmt.Println(match.Team2.Players[1].Info())

	var db psdb.PostgreDB
	db.InitDB()
}
