package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/uptrace/bun"
)

type Player struct {
	bun.BaseModel `bun:"table:players,alias:p"`
	ID            int64    `bun:"id,pk,autoincrement"`
	Nickname      string   `bun:"nickname,unique" json:"nickname"`
	Name          string   `bun:"name" json:"name"`
	Rank          Rank     `bun:"rank" json:"rank"`
	Description   string   `bun:"description" json:"description"`
	Age           int      `bun:"age" json:"age"`
	Position      Position `bun:"position" json:"position"`
	TotalGoals    int      `bun:"totalGoals" json:"totalGoals"`
	GamesPlayed   int      `bun:"gamesPlayed" json:"gamesPlayed"`
	GamesWon      int      `bun:"gamesWon" json:"gamesWon"`
	GamesLost     int      `bun:"gamesLost" json:"gamesLost"`
	Diff          int      `bun:"diff" json:"diff"`
	Elo           float64  `bun:"elo" json:"elo"`
	Points        float64  `bun:"points" json:"points"`
}

// Init creates the player with their values
func (p *Player) Init(nickname string, name string, description string, age int, rank Rank, position Position, totalGoals int, gamesPlayed int, gamesWon int, gamesLost int) error {

	if nickname == "" {
		msg := fmt.Sprint("nickname cannot be empty")
		log.Error().Msg(msg)
		return errors.New(msg)
	}
	if name == "" {
		msg := fmt.Sprint("name cannot be empty")
		log.Error().Msg(msg)
		return errors.New(msg)
	}
	if rank > Five || rank < One {
		msg := fmt.Sprintf("rank is not between 1 and 5: %d", rank)
		log.Error().Msg(msg)
		return errors.New(msg)
	}

	p.Nickname = nickname
	p.Name = name
	p.Description = description
	p.Rank = rank
	p.Position = position
	p.Age = age
	p.TotalGoals = totalGoals
	p.GamesPlayed = gamesPlayed
	p.GamesWon = gamesWon
	p.GamesLost = gamesLost
	p.Diff = p.GamesWon - p.GamesLost
	p.Elo = p.CalculateELO()
	p.Points = 0

	return nil
}

// CalculateELO calculates the player's ELO considering their points (age, goals, gamesWon, etc)
// is a simple formula, maybe improve it in a future
func (p *Player) CalculateELO() float64 {
	var ELO float64
	// the formula is rank * 5 - goals * 3 + gamesWon - (age-23) * 0.2
	ELO = float64(p.Rank*5) + float64(p.TotalGoals*3) + float64(p.Diff)*2.5 - float64(p.Age-23)*0.2
	return ELO
}

func (p *Player) ToJSON() []byte {
	bytes, _ := json.Marshal(p)
	return bytes
}
