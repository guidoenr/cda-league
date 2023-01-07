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
	Position      Position `bun:"position" json:"position"`
	GoalsPerMatch int      `bun:"goals" json:"goalsPerMatch"`
	GamesWon      int      `bun:"gamesWon" json:"gamesWon"`
	Age           int      `bun:"age" json:"age"`
	Elo           int      `bun:"elo" json:"elo"`
}

type JsonPlayer struct {
	Nickname      string `json:"nickname"`
	Name          string `json:"name"`
	Rank          int    `json:"rank"`
	Position      string `json:"position"`
	GoalsPerMatch int    `json:"goalsPerMatch"`
	GamesWon      int    `json:"gamesWon"`
	Age           int    `json:"age"`
}

// Init creates the player with their values
func (p *Player) Init(nickname string, name string, rank Rank, position Position, age int) error {

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
	p.Elo = int(p.Rank)*5 + p.GoalsPerMatch*3

	return nil
}

func (p *Player) Info() string {
	jsonData, _ := json.MarshalIndent(*p, "", "   ")
	return string(jsonData)
}

func (p *Player) updateRank(rank Rank) {
	// TODO db
	p.Rank = rank
}

func (p *Player) Show() {
	fmt.Printf("|%s| - %d - %s \n", p.Nickname, p.Rank, p.Position)
}
