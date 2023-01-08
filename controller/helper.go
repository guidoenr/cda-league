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

// CleanDB makes a DROP TABLE players;
func (h *Helper) CleanDB() error {
	_, err := h.db.BunDB.Query("DROP TABLE players;")
	if err != nil {
		msg := fmt.Sprintf("cleaning db: %v", err)
		return errors.New(msg)
	}
	return nil
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

// CreateSchemas create the 'player' table in the db
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
	// getting the bytes
	jsonData, _ := os.ReadFile("resources/players.json")

	// unmarshalling into JsonPlayers
	var jsonPlayers []model.Player
	err := json.Unmarshal(jsonData, &jsonPlayers)
	if err != nil {
		log.Error().Msgf("UNMARSHALLING")
	}

	// we must initialize each player to calculate the elo
	var newPlayer model.Player
	for _, p := range jsonPlayers {
		newPlayer.Init(p.Nickname, p.Name, p.Description, p.Age, p.Rank, p.Position, p.GoalsPerMatch, p.GamesWon)
		h.players = append(h.players, newPlayer)
	}

}
