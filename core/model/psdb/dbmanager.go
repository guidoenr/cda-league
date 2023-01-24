package psdb

import (
	"context"
	"encoding/json"
	"fmt"
	h "github.com/guidoenr/cda-league/core/handler"
	"github.com/guidoenr/cda-league/core/model"
	"github.com/rs/zerolog/log"
	"os"
	"time"
)

type DbManager struct {
	db *PostgreDB
}

func (dbm *DbManager) Init(db *PostgreDB) {
	dbm.db = db
}

// InitializeDatabase knows how to start the DB, checking the tables, making pings
// and doing all the operations of sanitizing the database
func (dbm *DbManager) InitializeDatabase(cleanDb ...bool) error {
	h.LogInfo("initializing database")

	// first let's ping the db
	err := dbm.pingToDB()
	if err != nil {
		msg := fmt.Sprintf("pinging to DB: %v", err)
		log.Warn().Msg(msg)
		log.Warn().Msg("waiting for 5 seconds...")
		time.Sleep(time.Second * 5)
		err = dbm.pingToDB()
		if err != nil {
			err := h.Newf(h.DbError, "unable to ping db: %v", err)
			return h.HandleError(err)
		}
	}

	// clean the DB
	if len(cleanDb) > 0 {
		err = dbm.cleanDB()
		if err != nil {
			err := h.Newf(h.DbError, "cleaning db: %v", err)
			return h.HandleError(err)
		}
	}

	// check if the tables were created
	if !dbm.tablesExist() {
		h.LogInfo("tables don't exist, creating tables...")
		// if the tables weren't created, create all the tables
		err = dbm.CreateTables()
		if err != nil {
			err := h.Newf(h.DbError, "creating tables: %v", err)
			return h.HandleError(err)
		}

		// now put the data inside
		h.LogInfo("tables created, dumping players/matches to DB")
		err = dbm.dumpPlayersToDB()
		if err != nil {
			err := h.Newf(h.DbError, "dumping players to db: %v", err)
			return h.HandleError(err)
		}
		h.LogInfo("tables created successfully")
	}

	// creating the backup (json files)
	// TODO, improve this?
	err = dbm.makeBackup()
	if err != nil {
		err := h.Newf(h.DbError, "making db backup: %v", err)
		return h.HandleError(err)
	}

	h.LogOk("database initialized successfully")
	return nil
}

// makeBackup will do a backup of the current state of the database in JSON files
// this backup is only in case the DB stop working
func (dbm *DbManager) makeBackup() error {
	// creating the variables to dump the data
	var players []model.Player
	var matches []model.Match

	// SELECT * FROM players
	err := dbm.db.BunDB.
		NewSelect().
		Model(&players).
		Scan(context.Background())

	// SELECT * FROM matches
	err = dbm.db.BunDB.
		NewSelect().
		Model(&matches).
		Scan(context.Background())

	if err != nil {
		err := h.Newf(h.DbError, "making backup: %v", err)
		return h.HandleError(err)
	}

	// marshall the structs into JSON
	JSONPlayers, _ := json.MarshalIndent(players, "", "   ")
	JSONMatches, _ := json.MarshalIndent(matches, "", "   ")

	// create the file names
	FilePlayers, err := dbm.generateBackupFileName("players")
	if err != nil {
		err := h.Newf(h.FileError, "creating backup filename: %v", err)
		return h.HandleError(err)
	}
	FileMatches, err := dbm.generateBackupFileName("matches")
	if err != nil {
		err := h.Newf(h.FileError, "creating backup filename: %v", err)
		return h.HandleError(err)
	}

	// write files
	err = os.WriteFile(FilePlayers, JSONPlayers, 0644)
	if err != nil {
		err := h.Newf(h.FileError, "writing file for players: %v", err)
		return h.HandleError(err)
	}
	err = os.WriteFile(FileMatches, JSONMatches, 0644)
	if err != nil {
		err := h.Newf(h.FileError, "writing file for matches: %v", err)
		return h.HandleError(err)
	}

	return nil
}

// generateBackupFileName returns the correct backup file name for the given table
func (dbm *DbManager) generateBackupFileName(tableName string) (string, error) {
	// get current date and time
	year, month, day := time.Now().Date()

	// creating the name of the file
	fileName := fmt.Sprintf("api/backup/{%d-%s-%d}-backup-%s.json", day, month.String(), year, tableName)

	_, err := os.Create(fileName)
	if err != nil {
		err := h.Newf(h.FileError, "writing file for matches: %v", err)
		return "", h.HandleError(err)
	}

	return fileName, nil
}

// pingToDB will check if the db is running
func (dbm *DbManager) pingToDB() error {
	err := dbm.db.BunDB.Ping()
	if err != nil {
		err := h.Newf(h.DbError, "making ping to db: %v", err)
		return h.HandleError(err)
	}
	return nil
}

// tablesExist will check if the db is running
func (dbm *DbManager) tablesExist() bool {
	var players []model.Player

	// SELECT * FROM players
	err := dbm.db.BunDB.
		NewSelect().
		Model(&players).
		Scan(context.Background())

	if err != nil || len(players) == 0 {
		return false
	}

	////TODO, think this
	// var matches []model.Match
	//// SELECT * FROM matches
	//err = dbm.db.BunDB.
	//	NewSelect().
	//	Model(&matches).
	//	Scan(context.Background())
	//
	//if err != nil || len(matches) == 0 {
	//	return false
	//}

	return true
}

// cleanDB makes a DROP TABLE players;
func (dbm *DbManager) cleanDB() error {
	_, err := dbm.db.BunDB.Query("DROP TABLE matches;")
	_, err = dbm.db.BunDB.Query("DROP TABLE players;")
	if err != nil {
		err := h.Newf(h.DbError, "droping tables: %v", err)
		return h.HandleError(err)
	}
	return nil
}

// dumpPlayersToDB initializes the db with all the players written in zero-day-players.json
func (dbm *DbManager) dumpPlayersToDB() error {
	// first read the json file
	players, err := dbm.readPlayersFromJSON()
	if err != nil {
		err := h.Newf(h.FileError, "reading players from JSON file: %v", err)
		return h.HandleError(err)
	}

	// making the query
	_, err = dbm.db.BunDB.NewInsert().
		Model(&players).
		Exec(context.Background())

	if err != nil {
		err := h.Newf(h.DbError, "dumping players to DB: %v", err)
		return h.HandleError(err)
	}
	return nil
}

// CreateTables create the 'player' table in the db TODO, add match later
func (dbm *DbManager) CreateTables() error {
	// creating players table
	_, err := dbm.db.BunDB.NewCreateTable().
		Model((*model.Player)(nil)).
		Exec(context.Background())

	if err != nil {
		err := h.Newf(h.DbError, "creating players table: %v", err)
		return h.HandleError(err)
	}

	_, err = dbm.db.BunDB.NewCreateTable().
		Model((*model.Match)(nil)).
		Exec(context.Background())

	if err != nil {
		err := h.Newf(h.DbError, "creating matches table: %v", err)
		return h.HandleError(err)
	}

	return nil
}

// readPlayersFromJSON reads all the players json list from /resources/zero-day-players.json
func (dbm *DbManager) readPlayersFromJSON() ([]model.Player, error) {
	var players []model.Player

	// reading the first players state
	jsonData, err := os.ReadFile("api/backup/zero-day-players.json")
	if err != nil {
		err := h.Newf(h.FileError, "reading zero-day-players.json: %v", err)
		return nil, h.HandleError(err)
	}

	// unmarshalling into JsonPlayers
	var jsonPlayers []model.Player
	err = json.Unmarshal(jsonData, &jsonPlayers)
	if err != nil {
		err := h.Newf(h.FileError, "unmarshalling players: %v", err)
		return nil, h.HandleError(err)
	}

	// we must initialize each player to calculate the elo
	var newPlayer model.Player
	for _, p := range jsonPlayers {
		err = newPlayer.Init(p.Nickname, p.Name, p.Description, p.Age, p.Rank, p.Position, p.TotalGoals, p.GamesPlayed, p.GamesWon, p.GamesLost)
		if err != nil {
			err := h.Newf(h.PlayerError, "initialazing player: %v", err)
			return nil, h.HandleError(err)
		}
		players = append(players, newPlayer)
	}

	return players, nil
}
