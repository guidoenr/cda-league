package model

// Result of a match
type Result struct {
	PlayerGoals map[string]int `json:"playersGoals"` // the goals that each player made (e.g: {1 : 3} - where 1 references: "Messi" that scored 3 goals)
	Winner      Team           `json:"winner"`       // winner of the match
	Loser       Team           `json:"loser"`        // loser of the match
}

// ThePlayerWon check if the player is in the winner team or not
// TODO, improve this. Time complexity of O(n) but the maximum n will be 7
func (r *Result) ThePlayerWon(player Player) bool {
	for _, p := range r.Winner.Players {
		if p.ID == player.ID {
			return true
		}
	}
	return false
}
