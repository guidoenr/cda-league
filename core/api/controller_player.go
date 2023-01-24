package api

import (
	"context"
	"fmt"
	h "github.com/guidoenr/cda-league/core/handler"
	"github.com/guidoenr/cda-league/core/model"
	"github.com/guidoenr/cda-league/core/model/psdb"
	"github.com/rs/zerolog/log"
	"strconv"
)

// ControllerPlayer handles all the player persist logic
type ControllerPlayer struct {
	db *psdb.PostgreDB
}

func (cp *ControllerPlayer) Init(db *psdb.PostgreDB) {
	cp.db = db
}

// GetPlayers (GET /players)
// check in the DB all the stored players and return them in JSON format
func (cp *ControllerPlayer) GetPlayers() ([]model.Player, error) {
	log.Info().Msgf("getting players")
	var players []model.Player

	// SELECT * FROM players
	err := cp.db.BunDB.
		NewSelect().
		Model(&players).
		Scan(context.Background())

	if err != nil {
		err := h.Newf(h.PlayerError, "getting players from DB: %v", err)
		return nil, h.HandleError(err)
	}

	return players, nil
}

// GetPlayersRankedBy : sortField (might be elo, rank, age) | sortFields (optional) - list of more fields to sort
// (GET /players)
// returns all the players sorted by their elo, maybe you can think
// "ok, but you can use getplayers and then sort it" but no...
// queries in an SQL motor are more performant than in the code
func (cp *ControllerPlayer) GetPlayersRankedBy(sortField string, sortFields ...string) ([]model.Player, error) {
	log.Info().Msgf("getting ranked by elo players")
	var players []model.Player
	var orderExpr string

	// generating the orderExpr
	orderExpr = fmt.Sprintf("%s DESC", sortField)
	for _, field := range sortFields {
		orderExpr = fmt.Sprintf("%s, %s DESC", orderExpr, field)
	}

	// SELECT * FROM players ORDER BY $sortField
	err := cp.db.BunDB.
		NewSelect().
		Model(&players).
		OrderExpr(orderExpr).
		Scan(context.Background())

	// WARNING with this
	if err != nil {
		err := h.Newf(h.PlayerError, "getting players rankedBy-> %s %s: %v", sortField, sortFields, err)
		return nil, h.HandleError(err)
	}

	return players, nil
}

// GetPlayerByID (GET /players/:id)
// check in the DB the player finding by id
func (cp *ControllerPlayer) GetPlayerByID(id string) (model.Player, error) {
	var player model.Player
	playerId, _ := strconv.Atoi(id)

	// find the player by id
	err := cp.db.BunDB.
		NewSelect().
		Model(&player).
		Where("id = ?", playerId).
		Scan(context.Background())

	if err != nil {
		err := h.Newf(h.PlayerError, "getting player by id '%s': %v", id, err)
		return player, h.HandleError(err)
	}

	return player, nil
}

// GetPlayerByNickname (GET /players/) found a player by their nickname
func (cp *ControllerPlayer) GetPlayerByNickname(nickname string) (model.Player, error) {
	var player model.Player

	// find player by nickname
	err := cp.db.BunDB.
		NewSelect().
		Model(&player).
		Where("nickname = ?", nickname).
		Scan(context.Background())

	if err != nil {
		err := h.Newf(h.PlayerError, "getting player by nickname '%s': %v", nickname, err)
		return player, h.HandleError(err)
	}

	return player, nil
}

// UpdatePlayers (POST /players/:id/update) updates the players fields given a match result
// and return the updated rows count
// (NOTE: this function receives a map[int]int that contains the id of the player and the goals
func (cp *ControllerPlayer) UpdatePlayers(matchResult model.Result) (int64, error) {

	// validate player IDs
	for playerId := range matchResult.PlayerGoals {
		if _, err := cp.GetPlayerByID(playerId); err != nil {
			err := h.Newf(h.PlayerError, "invalid player id '%s': %v", playerId, err)
			return -1, h.HandleError(err)
		}
	}

	// create the players slice to update
	players := make([]model.Player, 0, len(matchResult.PlayerGoals))

	for playerId, goals := range matchResult.PlayerGoals {
		// get the player data
		player, err := cp.GetPlayerByID(playerId)
		if err != nil {
			err := h.Newf(h.PlayerError, "getting player ids to update: %v", err)
			return -1, h.HandleError(err)
		}

		// check if the player won
		won := matchResult.ThePlayerWon(player)

		// update the player stats
		player.UpdatePlayer(goals, won)
		players = append(players, player)
	}

	// update all players in a batch
	result, err := cp.db.BunDB.
		NewUpdate().
		Model(&players).
		Exec(context.Background())

	if err != nil {
		err := h.Newf(h.PlayerError, "updating players by match result: %v", err)
		return -1, h.HandleError(err)
	}

	return result.RowsAffected()
}
