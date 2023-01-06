package model

import (
	"encoding/json"
	"fmt"
	"os"
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

	if name == "aux" {
		jsonData, _ := json.MarshalIndent(t, "", "   ")
		os.WriteFile("temp-aux.json", jsonData, 0644)
	}
}

// goodMixWith determines if the teams that the match mixed are a good shuffle
// this method is going to be used only for the 'auxTeam'
func (t *Team) goodMixWith(otherTeam *Team) bool {
	bytes, _ := os.ReadFile("temp-aux.json")

	var auxTeam Team
	json.Unmarshal(bytes, &auxTeam)

	fmt.Printf("----------------AUX TEAM\n")
	auxTeam.Show()
	fmt.Printf("----------------AUX TEAM\n")

	var count int
	for _, playerOne := range auxTeam.Players {
		for _, playerTwo := range otherTeam.Players {
			if playerTwo.Nickname == playerOne.Nickname {
				count++
			}
		}
	}
	fmt.Println(count)
	return count < 3
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
