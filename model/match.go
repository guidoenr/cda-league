package model

import (
	"github.com/guidoenr/cda-league/config"
	"github.com/uptrace/bun"
	"math"
	"math/rand"
	"time"
)

type Match struct {
	bun.BaseModel `bun:"table:matches,alias:m"`
	ID            int64             `bun:"id,pk,autoincrement"`
	PlayersByRank map[Rank][]Player // the players ordered by their rank
	Players       []Player          `json:"players"`
	Date          time.Time         `bun:"date"   json:"date"`   // the date of the team
	Field         string            `bun:"field"  json:"field"`  // field in where the match was played
	Team1         Team              `bun:"team1"  json:"team1"`  // the first team
	Team2         Team              `bun:"team2"  json:"team2"`  // the second team
	MatchResult   Result            `bun:"result" json:"result"` // the match result
}

// GenerateTeams is the complete algorithm to create evenly teams based on the player's ranking
// where the first picker is random and each player are distributed based on their ranks
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
		// group is the list of players of a rank, example rank 4 : [player1, player2, .. , playerN]
		rand.Shuffle(len(group), func(i, j int) {
			group[i], group[j] = group[j], group[i]
		})
		playersByRank[i] = group
	}

	// setting the Players by rank
	m.PlayersByRank = playersByRank

	// creating the two teams map
	// teams = {
	// 0: [p1, p2, p3, p4, p5] --> team 1
	// 1: [j1, j2, j3, j4, j5] --> team 2
	teams := make([][]Player, 2)

	// choosing the first/second picker
	// if the first was "0" then the second will be "-1"
	// and with math.Abs will be "1"
	first := rand.Intn(2)
	second := int(math.Abs(float64(first - 1)))

	// distributing the players,
	rank := Five
	for rank > 0 {
		for _, p := range m.PlayersByRank[rank] {
			if len(teams[first]) > len(teams[second]) {
				teams[second] = append(teams[second], p)
			} else {
				teams[first] = append(teams[first], p)
			}
		}
		rank--
	}

	// creating the teams
	var team1, team2 Team
	team1Name, team2Name := m.pickTwoTeamNames()

	team1.Init(team1Name, len(teams[0]), teams[0])
	team2.Init(team2Name, len(teams[1]), teams[1])

	// calculate the total points
	totalPoints := team1.Points + team2.Points

	// setting the chance of winning of each map
	team1.setChanceOfWinning(totalPoints)
	team2.setChanceOfWinning(totalPoints)

	m.Team1 = team1
	m.Team2 = team2
}

// pickTwoTeamNames search for the lisf of teams names and chose two random team names
func (m *Match) pickTwoTeamNames() (string, string) {
	// reading the config
	var cfg config.Config
	cfg.Init()
	teamsNames := cfg.TeamNames

	size := len(teamsNames)

	num1, num2 := rand.Intn(size), rand.Intn(size)

	// Make sure the numbers are different
	for num1 == num2 {
		num2 = rand.Intn(len(teamsNames))
	}

	return teamsNames[num1], teamsNames[num2]
}
