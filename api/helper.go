package api

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/guidoenr/cda-league/handler"
	"github.com/guidoenr/cda-league/model"
	"github.com/guidoenr/cda-league/model/psdb"
	"github.com/rs/zerolog/log"
	"os"
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
func (h *Helper) InitializeDatabase(cleanDb ...bool) error {
	log.Info().Msg("initializing database...")

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

	// clean the DB
	if len(cleanDb) > 0 {
		err = h.CleanDB()
		if err != nil {
			return handler.HandleError("cleaning the DB: %v", err)
		}
	}

	// check if the tables were created
	if !h.TablesCreated() {
		log.Info().Msg("tables don't exist, creating tables...")
		// if the tables weren't created, create all the tables
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
		log.Info().Msg("tables created successfully")
	}

	// creating the backup (json files)
	// TODO, improve this?
	err = h.MakeBackup()
	if err != nil {
		return handler.HandleError("making backup: %v", err)
	}

	log.Info().Msg("[Database OK] - Ready to go ;)")
	return nil
}

// MakeBackup will do a backup of the current state of the database in JSON files
// this backup is only in case the DB stop working
func (h *Helper) MakeBackup() error {
	// creating the variables to dump the data
	var players []model.Player
	var matches []model.Match

	// SELECT * FROM players
	err := h.db.BunDB.
		NewSelect().
		Model(&players).
		Scan(context.Background())

	// SELECT * FROM matches
	err = h.db.BunDB.
		NewSelect().
		Model(&matches).
		Scan(context.Background())

	if err != nil {
		return handler.HandleError("[critical]- making backup: %v", err)
	}

	// marshall the structs into JSON
	JSONPlayers, _ := json.MarshalIndent(players, "", " ")
	JSONMatches, _ := json.MarshalIndent(matches, "", " ")

	// create the file names
	FilePlayers, err := h.generateBackupFileName("players")
	if err != nil {
		return handler.HandleError("creating backup filename for players: %v", err)
	}
	FileMatches, err := h.generateBackupFileName("matches")
	if err != nil {
		return handler.HandleError("creating backup filename for matches: %v", err)
	}

	// write files
	err = os.WriteFile(FilePlayers, JSONPlayers, 0644)
	if err != nil {
		return handler.HandleError("writing file for players: %v", err)
	}
	err = os.WriteFile(FileMatches, JSONMatches, 0644)
	if err != nil {
		return handler.HandleError("writing file for matches: %v", err)
	}

	if err != nil {
		return handler.HandleError("[critical]- dumping players into backup files: %v", err)
	}

	return nil
}

// generateBackupFileName returns the correct backup file name for the given table
func (h *Helper) generateBackupFileName(tableName string) (string, error) {
	// get current date and time
	year, month, day := time.Now().Date()
	hour := time.Now().Hour()
	minute := time.Now().Minute()

	// creating the name of the file
	fileName := fmt.Sprintf("api/backup/{%d-%s-%d|%d:%d}-backup-%s.json", day, month.String(), year, hour, minute, tableName)

	_, err := os.Create(fileName)
	if err != nil {
		return "", handler.HandleError("creating filename '%s': %v", fileName, err)
	}

	return fileName, nil
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
	return true
}

// CleanDB makes a DROP TABLE players;
func (h *Helper) CleanDB() error {
	_, err := h.db.BunDB.Query("DROP TABLE matches;")
	_, err = h.db.BunDB.Query("DROP TABLE players;")
	if err != nil {
		msg := fmt.Sprintf("cleaning db: %v", err)
		return errors.New(msg)
	}
	return nil
}

// DumpPlayersToDB initializes the db with all the players written in zero-day-players.json
func (h *Helper) DumpPlayersToDB() error {
	// first read the json file
	err := h.readPlayersFromJSON()
	if err != nil {
		return handler.HandleError("reading players from json: %v", err)
	}

	// making the query
	_, err = h.db.BunDB.NewInsert().
		Model(&h.players).
		Exec(context.Background())

	if err != nil {
		return handler.HandleError("dumping players to db: %v", err)
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
		return handler.HandleError("creating players table: %v", err)
	}

	_, err = h.db.BunDB.NewCreateTable().
		Model((*model.Match)(nil)).
		Exec(context.Background())

	if err != nil {
		return handler.HandleError("creating matches table: %v", err)
	}

	return nil
}

// readPlayersFromJSON reads all the players json list from /resources/zero-day-players.json
func (h *Helper) readPlayersFromJSON() error {
	// reading the first players state
	jsonData, err := os.ReadFile("api/backup/zero-day-players.json")
	if err != nil {
		return handler.HandleError("reading zero day players file : %v", err)
	}

	// unmarshalling into JsonPlayers
	var jsonPlayers []model.Player
	err = json.Unmarshal(jsonData, &jsonPlayers)
	if err != nil {
		return handler.HandleError("unmarshalling players: %v", err)
	}

	// we must initialize each player to calculate the elo
	var newPlayer model.Player
	for _, p := range jsonPlayers {
		newPlayer.Init(p.Nickname, p.Name, p.Description, p.Age, p.Rank, p.Position, p.TotalGoals, p.GamesPlayed, p.GamesWon, p.GamesLost)
		h.players = append(h.players, newPlayer)
	}

	return nil
}
