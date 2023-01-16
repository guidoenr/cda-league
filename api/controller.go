package api

import (
	"context"
	"errors"
	"fmt"
	"github.com/guidoenr/fulbo/model"
	"github.com/guidoenr/fulbo/model/psdb"
	"github.com/rs/zerolog/log"
	"strconv"
)

type PlayerControler struct {
	db *psdb.PostgreDB
}

func (pc *PlayerControler) Init(db *psdb.PostgreDB) {
	pc.db = db
}

// GetPlayers (GET /players)
// check in the DB all the stored players and return them in JSON format
func (pc *PlayerControler) GetPlayers() ([]model.Player, error) {
	log.Info().Msgf("getting players")
	var players []model.Player

	// SELECT * FROM players
	err := pc.db.BunDB.NewSelect().
		Model(&players).
		Scan(context.Background())

	if err != nil {
		msg := fmt.Sprintf("getting all players - select query to db: %v", err)
		log.Error().Msg(msg)
		return players, errors.New(msg)
	}

	return players, nil
}

// GetPlayersRankedBy : sortField (might be elo, rank, age) | sortFields (aditionals)
// (GET /players)
// returns all the players sorted by their elo, maybe you can think
// "ok, but you can use getplayers and then sort it" but no...
// queries in an SQL motor are more performant than in the code
func (pc *PlayerControler) GetPlayersRankedBy(sortField string, sortFields ...string) ([]model.Player, error) {
	log.Info().Msgf("getting ranked by elo players")
	var players []model.Player
	var orderExpr string

	// generating the orderExpr
	orderExpr = fmt.Sprintf("%s DESC", sortField)
	for _, field := range sortFields {
		orderExpr = fmt.Sprintf("%s, %s DESC", orderExpr, field)
	}

	// SELECT * FROM players ORDER BY $sortField
	err := pc.db.BunDB.NewSelect().
		Model(&players).
		OrderExpr(orderExpr).
		Scan(context.Background())

	if err != nil {
		msg := fmt.Sprintf("getting all players ranked by %s - select query to db: %v", sortFields, err)
		log.Error().Msg(msg)
		return players, errors.New(msg)
	}

	return players, nil
}

// GetPlayerByID (GET /players/:id)
// check in the DB the player finding by id
func (pc *PlayerControler) GetPlayerByID(id string) (model.Player, error) {
	var player model.Player
	playerId, _ := strconv.Atoi(id)

	err := pc.db.BunDB.NewSelect().
		Model(&player).
		Where("id = ?", playerId).
		Scan(context.Background())

	if err != nil {
		msg := fmt.Sprintf("selecting player '%s' by id: %v", id, err)
		log.Error().Msg(msg)
		return player, errors.New(msg)
	}

	return player, nil
}

// GetPlayerByNickname (GET /players/) found a player by their nickname
func (pc *PlayerControler) GetPlayerByNickname(nickname string) (model.Player, error) {
	var player model.Player

	err := pc.db.BunDB.NewSelect().
		Model(&player).
		Where("nickname = ?", nickname).
		Scan(context.Background())

	if err != nil {
		msg := fmt.Sprintf("selecting player '%s' by nickname: %v", nickname, err)
		log.Error().Msg(msg)
		return player, errors.New(msg)
	}

	return player, nil
}

// UpdatePlayer (POST /players/:id/update) updates the player fields
func (pc *PlayerControler) UpdatePlayer() (string, error) {
	/*var player model.Player

	err, _ := db.BunDB.NewUpdate().
		Model(&player).
		Where("id = ?", id).
		Set("nickname = ?", nickname).
		Set("age = ?", age).
		Set("rank = ?", rank).
		Set("position = ?", position).
		Set("goalsPerMatch = ?", goalsPerMatch).
		Set("gamesWon = ?", gamesWon).
		Exec(context.Background())

	if err != nil {
		msg := fmt.Sprintf("updating player fields '%s' by id: %v", id, err)
		log.Error().Msg(msg)
		return player, errors.New(msg)
	}*/
	return "TODO--> NOT IMPLEMENTED YET", nil
}

func (pc *PlayerControler) CreatePlayer() (string, error) {
	var players string

	return players, nil
}
