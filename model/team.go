package model

import (
	"encoding/json"
	"fmt"
)

type Team struct {
	Name            string   `json:"name"`
	TotalPlayers    int      `json:"totalPlayers"`
	Players         []Player `json:"players"`
	points          float64  `json:"points"`
	ChanceOfWinning string   `json:"chanceOfWinning"`
}

func (t *Team) Init(name string, totalPlayers int, players []Player) {
	t.Name = name
	t.TotalPlayers = totalPlayers
	t.Players = players

	// the total points of a team is the sum of the Elo's of each player
	for _, p := range players {
		t.points += p.Elo
	}
}

// goodMixWith determines if the teams that the match mixed are a good shuffle
// this method is going to be used only for the 'auxTeam'
func (t *Team) goodMixWith(otherTeam *Team) bool {
	// TODO THIS
	return true
}

func (t *Team) setChanceOfWinning(totalPoints float64) {
	t.ChanceOfWinning = fmt.Sprintf("%.2f", (t.points/totalPoints)*100)
}

func (t *Team) ToJSON() string {
	bytes, _ := json.Marshal(t)
	return string(bytes)
}
