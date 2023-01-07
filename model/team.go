package model

import "encoding/json"

type Team struct {
	Name            string
	TotalPlayers    int
	Players         []Player
	Points          float64
	ChanceOfWinning float64
}

func (t *Team) Init(name string, totalPlayers int, players []Player) {
	t.Name = name
	t.TotalPlayers = totalPlayers
	t.Players = players

	for _, p := range players {
		t.Points += float64(p.Elo)
	}
}

// goodMixWith determines if the teams that the match mixed are a good shuffle
// this method is going to be used only for the 'auxTeam'
func (t *Team) goodMixWith(otherTeam *Team) bool {
	// TODO THIS
	return true
}

func (t *Team) setChanceOfWinning(totalPoints float64) {
	t.ChanceOfWinning = (t.Points / totalPoints) * 100
}

func (t *Team) ToJSON() string {
	bytes, _ := json.Marshal(t)
	return string(bytes)
}
