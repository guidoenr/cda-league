package controller

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/guidoenr/fulbo/model"
	"github.com/guidoenr/fulbo/model/psdb"
	"github.com/rs/zerolog/log"
	"os"
)

type Helper struct {
	db      *psdb.PostgreDB
	players []model.Player
}

func (h *Helper) Init(db *psdb.PostgreDB) {
	h.db = db
}

// DumpPlayersToDB initializes the db with all the players written in players.json
func (h *Helper) DumpPlayersToDB() error {
	// first read the json file
	h.readPlayersFromJSON()

	// making the query
	_, err := h.db.BunDB.NewInsert().
		Model(&h.players).
		Exec(context.Background())

	if err != nil {
		msg := fmt.Sprintf("dumping players to db: %v", err)
		log.Error().Msg(msg)
		return errors.New(msg)
	}
	return nil
}

func (h *Helper) CreateSchemas() error {
	// creating players table
	_, err := h.db.BunDB.NewCreateTable().
		Model((*model.Player)(nil)).
		Exec(context.Background())

	if err != nil {
		msg := fmt.Sprintf("creating schemas: %v", err)
		log.Error().Msg(msg)
		return errors.New(msg)
	}
	return nil
}

// readPlayersFromJSON reads all the players json list from /resources/players.json
func (h *Helper) readPlayersFromJSON() {
	var players []model.Player
	// getting the bytes
	jsonData, _ := os.ReadFile("resources/players.json")

	// unmarshalling into JsonPlayers
	var jsonPlayers []model.JsonPlayer
	json.Unmarshal(jsonData, &jsonPlayers)

	// now parsing
	var player model.Player
	for _, p := range jsonPlayers {
		player.Init(p.Nickname, p.Name, model.Rank(p.Rank), model.Position(p.Position), p.Age)
		players = append(players, player)
	}

	h.players = players
}
