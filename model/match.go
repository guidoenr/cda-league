package model

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type Match struct {
	Players       []Player `json:"Players"`
	PlayersByRank map[Rank][]Player
	Team1         Team
	Team2         Team
	Winner        Team
}

func (m *Match) Init(players []Player) {
	m.Players = players
}

// InitFromJson set the players into the match, read from /resources/players.json
func (m *Match) InitFromJson() error {
	jsonData, err := os.ReadFile("resources/players.json")
	if err != nil {
		return fmt.Errorf("error reading file")
	}

	var players []JsonPlayer
	err = json.Unmarshal(jsonData, &players)
	if err != nil {
		return fmt.Errorf("error unmarshalling: %v", err)
	}

	var player Player
	for _, p := range players {
		_ = player.Init(p.Nickname, p.Name, Rank(p.Rank), Position(p.Position), p.Age)
		m.Players = append(m.Players, player)
	}

	return nil
}

// GenerateTeams is the complete algorithm to create evenly teams based on the player's ranking
func (m *Match) GenerateTeams() {
	// the Players by rank have a structure like the following
	// 5 : [ Messi, Maradona ]
	// 4 : [ Mbappe, Cristiano]
	playersByRank := make(map[Rank][]Player)

	// search for each player rank and append it to the rank slice
	for _, player := range m.Players {
		rank := player.Rank
		// appending each player to their rank
		playersByRank[rank] = append(playersByRank[rank], player)
	}

	// now shuffle the teams using the nano seconds as seed
	rand.Seed(time.Now().UnixNano())
	for i, group := range playersByRank {
		rand.Shuffle(len(group), func(i, j int) {
			group[i], group[j] = group[j], group[i]
		})
		playersByRank[i] = group
	}

	// setting the Players by rank
	m.PlayersByRank = playersByRank
	players1, players2 := make([]Player, 0), make([]Player, 0)

	i := 5
	for i > 0 {
		for _, p := range m.PlayersByRank[Rank(i)] {
			if len(players1) > len(players2) {
				players2 = append(players2, p)
			} else {
				players1 = append(players1, p)
			}
		}
		i--
	}

	// this aux team is only to compare the players of a team
	var auxTeam Team
	auxTeam.init("aux", len(players1), players1)

	// creating the teams
	var team1, team2 Team

	team1.init("Team1", len(players1), players1)
	team2.init("Team2", len(players2), players2)

	// if the teams are some kind of 'equal' (3 players are the same)
	// then we start the algorithm again
	if !auxTeam.goodMixWith(&team1) {
		fmt.Printf("------------------------------ REPEATED")
		m.GenerateTeams()
	}

	// calculate the total points
	totalPoints := team1.Points + team2.Points

	team1.setChanceOfWinning(totalPoints)
	team2.setChanceOfWinning(totalPoints)

	m.Team1 = team1
	m.Team2 = team2
}
