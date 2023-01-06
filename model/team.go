package model

import (
	"fmt"
)

type Team struct {
	Name            string
	TotalPlayers    int
	Players         []Player
	Points          float64
	ChanceOfWinning float64
}

func (t *Team) init(name string, totalPlayers int, players []Player) {
	t.Name = name
	t.TotalPlayers = totalPlayers
	t.Players = players

	for _, p := range players {
		t.Points += float64(p.Elo)
	}

}

func (t *Team) setChanceOfWinning(score float64) {
	t.ChanceOfWinning = score
}

func (t *Team) Show() {
	fmt.Printf("-----[Team: %s]-----\n", t.Name)
	fmt.Println("-----------------------")
	for _, p := range t.Players {
		p.Show()
	}
	fmt.Println()
	fmt.Printf("chance of winning: %.2f \n", t.ChanceOfWinning*100)
}
