package model

import (
	"math/rand"
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

	// now shuffle the teams using the nanoseconds as seed
	rand.Seed(time.Now().UnixNano())
	for i, group := range playersByRank {
		rand.Shuffle(len(group), func(i, j int) {
			group[i], group[j] = group[j], group[i]
		})
		playersByRank[i] = group
	}

	// setting the Players by rank
	m.PlayersByRank = playersByRank

	// creating the two teams
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

	// creating the teams
	var team1, team2 Team

	team1.Init("Team1", len(players1), players1)
	team2.Init("Team2", len(players2), players2)

	/*
		// this aux team is only to compare the players of a team
		var auxTeam Team
		auxTeam.Init("aux", len(players1), players1)
		// if the teams are some kind of 'equal' (3 players are the same)
		// then we start the algorithm again
		if !auxTeam.goodMixWith(&team1) {
			m.GenerateTeams()
			return
		} */

	// calculate the total points
	totalPoints := team1.Points + team2.Points

	team1.setChanceOfWinning(totalPoints)
	team2.setChanceOfWinning(totalPoints)

	m.Team1 = team1
	m.Team2 = team2
}
