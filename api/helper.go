package api

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/guidoenr/fulbo/handler"
	"github.com/guidoenr/fulbo/model"
	"github.com/guidoenr/fulbo/model/psdb"
	"github.com/rs/zerolog/log"
	"os"
	"strings"
	"time"
)

type Helper struct {
	db      *psdb.PostgreDB
	players []model.Player
}

func (h *Helper) Init(db *psdb.PostgreDB) {
	h.db = db
}

// InitializeDatabase knows how to start the DB, checking the tables, making pings
// and doing all the operations of sanitizing the database
func (h *Helper) InitializeDatabase() error {
	// first let's ping the db
	err := h.PingToDb()
	if err != nil {
		msg := fmt.Sprintf("pinging to DB: %v", err)
		log.Warn().Msg(msg)
		log.Warn().Msg("waiting for 5 seconds...")
		time.Sleep(time.Second * 5)
		err = h.PingToDb()
		if err != nil {
			return handler.HandleError("unable to ping to DB: %v", err)
		}
	}

	// check if the tables were created
	tablesCreated := h.TablesCreated()
	if !tablesCreated {
		err = h.CreateTables()
		if err != nil {
			return handler.HandleError("creating tables: %v", err)
		}

		// now put the data inside
		log.Info().Msg("tables created, filling the DB..")
		err = h.DumpPlayersToDB()
		if err != nil {
			return handler.HandleError("dumping players to db: %v", err)
		}
		log.Info().Msg("tables created succesfully")
	}

	log.Info().Msg("[Database OK] - Ready to go ;)")
	return nil
}

// PingToDb will check if the db is running
func (h *Helper) PingToDb() error {
	err := h.db.BunDB.Ping()
	if err != nil {
		msg := fmt.Sprintf("making ping to DB: %v", err)
		log.Error().Msg(msg)
		return errors.New(msg)
	}
	return nil
}

// TablesCreated will check if the db is running
func (h *Helper) TablesCreated() bool {
	created := false

	// trying to create the tables
	err := h.CreateTables()
	if strings.Contains(err.Error(), "exists") {
		created = true
	}

	return created
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
	// TODO, maybe create a backup?
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

// CreateTables create the 'player' table in the db TODO, add match later
func (h *Helper) CreateTables() error {
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
		newPlayer.Init(p.Nickname, p.Name, p.Description, p.Age, p.Rank, p.Position, p.TotalGoals, p.GamesPlayed, p.GamesWon, p.GamesLost)
		h.players = append(h.players, newPlayer)
	}

}
