package controller

import (
	"context"
	"errors"
	"fmt"
	"github.com/guidoenr/fulbo/model"
	"github.com/guidoenr/fulbo/model/psdb"
	"github.com/rs/zerolog/log"
	"strconv"
)

// GetPlayers (GET /players)
// check in the DB all the stored players and return them in JSON format
func GetPlayers(db *psdb.PostgreDB) ([]model.Player, error) {
	log.Info().Msgf("getting players")
	var players []model.Player

	// SELECT * FROM players
	err := db.BunDB.NewSelect().
		Model(&players).
		Scan(context.Background())

	if err != nil {
		msg := fmt.Sprintf("getting all players - select query to db: %v", err)
		log.Error().Msg(msg)
		return players, errors.New(msg)
	}

	return players, nil
}

// GetPlayerByID (GET /players/:id)
// check in the DB the player finding by id
func GetPlayerByID(db *psdb.PostgreDB, id string) (model.Player, error) {
	var player model.Player
	playerId, _ := strconv.Atoi(id)

	err := db.BunDB.NewSelect().
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
func GetPlayerByNickname(db *psdb.PostgreDB, nickname string) (model.Player, error) {
	var player model.Player

	err := db.BunDB.NewSelect().
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
func UpdatePlayer(db *psdb.PostgreDB) (string, error) {
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

func CreatePlayer(db *psdb.PostgreDB) (string, error) {
	var players string

	return players, nil
}
