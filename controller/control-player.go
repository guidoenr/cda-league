package controller

import (
	"context"
	"errors"
	"github.com/guidoenr/fulbo/model"
	"github.com/guidoenr/fulbo/model/psdb"
	"github.com/rs/zerolog/log"
	"strconv"
)

// GetPlayers (GET /players)
// check in the DB all the stored players and return them in JSON format
func GetPlayers(db *psdb.PostgreDB) (string, error) {
	log.Info().Msgf("getting players")
	var players []model.Player

	// making the query
	err := db.BunDB.NewSelect().
		Model(&players).
		Scan(context.Background())

	if err != nil {
		return "", errors.New(err.Error())
	}

	playersString := ""
	for _, p := range players {
		playersString += p.Info()
	}

	return playersString, nil
}

// GetPlayerByID (GET /players/:id)
// check in the DB the player finding by id
func GetPlayerByID(db *psdb.PostgreDB, id string) (string, error) {
	var player model.Player
	playerId, _ := strconv.Atoi(id)

	err := db.BunDB.NewSelect().
		Model(&player).
		Where("id = ?", playerId).
		Scan(context.Background())

	if err != nil {
		return "", errors.New(err.Error())
	}

	return player.Info(), nil
}

func GetPlayerByNickname(db *psdb.PostgreDB, nickname string) (string, error) {
	var player model.Player

	err := db.BunDB.NewSelect().
		Model(&player).
		Where("nickname = ?", nickname).
		Scan(context.Background())

	if err != nil {
		return "", errors.New(err.Error())
	}

	return player.Info(), nil
}

func CreatePlayer(db *psdb.PostgreDB) (string, error) {
	var players string

	return players, nil
}
