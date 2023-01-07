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
	Age           int      `bun:"age" json:"age"`
	Rank          Rank     `bun:"rank" json:"rank"`
	Position      Position `bun:"position" json:"position"`
	Elo           int      `bun:"elo" json:"elo"`
	GoalsPerMatch int      `bun:"goalsPerMatch" json:"goalsPerMatch"`
	GamesWon      int      `bun:"gamesWon" json:"gamesWon"`
}

// Init creates the player with their values
func (p *Player) Init(nickname string, name string, age int, rank Rank, position Position, goalsPerMatch int, gamesWon int) error {

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
	p.Rank = rank
	p.Position = position
	p.Age = age
	p.GoalsPerMatch = goalsPerMatch
	p.GamesWon = gamesWon
	p.Elo = int(p.Rank)*5 + p.GoalsPerMatch*3 + p.GamesWon

	return nil
}

func (p *Player) ToJSON() []byte {
	bytes, _ := json.Marshal(p)
	return bytes
}
