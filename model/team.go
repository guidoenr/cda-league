package model

import (
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog/log"
	"os"
)

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

	// saving the team as a file
	fileName := fmt.Sprintf("%s.json", t.Name)
	jsonData, _ := json.MarshalIndent(t, "", "   ")
	err := os.WriteFile(fileName, jsonData, 0644)
	if err != nil {
		log.Error().Msgf("writing file: %v", err)
	}
}

// goodMixWith determines if the teams that the match mixed are a good shuffle
// this method is going to be used only for the 'auxTeam'
func (t *Team) goodMixWith(otherTeam *Team) bool {
	// TODO THIS
	return true
}

func (t *Team) setChanceOfWinning(totalPoints float64) {
	t.ChanceOfWinning = t.Points / totalPoints
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
