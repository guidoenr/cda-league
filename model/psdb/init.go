package psdb

import (
	"crypto/tls"
	"database/sql"
	"errors"
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"os"
)

// PostgreDB structs contains the bunDB database (initialized) to interact with the DB
type PostgreDB struct {
	BunDB     *bun.DB
	dbname    string
	connector *pgdriver.Connector
}

// InitDB initializes the postgresql database
func (pdb *PostgreDB) InitDB() error {
	// loading the connector
	pdb.loadConnector()

	// opening the DB
	sqldb := sql.OpenDB(pdb.connector)

	// creating new bun DB
	bunDB := bun.NewDB(sqldb, pgdialect.New())

	// connection pool
	bunDB.SetMaxIdleConns(5)
	bunDB.SetMaxOpenConns(10)

	// pinging to DB
	err := bunDB.Ping()
	if err != nil {
		msg := fmt.Sprintf("pinging to bunDB: %v", err)
		log.Error().Msgf(msg)
		return errors.New(msg)
	}

	// setting the bunDB
	pdb.BunDB = bunDB

	log.Info().Msgf("connected to bunDB '%s' succesfully", pdb.dbname)
	return nil
}

// CloseDB closes the postgresql database
func (pdb *PostgreDB) CloseDB() error {
	err := pdb.BunDB.Close()
	if err != nil {
		msg := fmt.Sprintf("closing db: %v", err)
		log.Error().Msgf(msg)
		return errors.New(msg)
	}

	log.Info().Msg("database closed successfully")
	return nil
}

// loadConnector returns the connector with the ENV variables to connect to the DB
func (pdb *PostgreDB) loadConnector() {
	dbname := os.Getenv("DB_NAME")
	network := os.Getenv("DB_NETWORK")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	addr := os.Getenv("DB_ADDR")

	// initializing the connector
	pgconn := pgdriver.NewConnector(
		pgdriver.WithNetwork(network),
		pgdriver.WithAddr(addr),
		pgdriver.WithInsecure(true),
		pgdriver.WithTLSConfig(&tls.Config{InsecureSkipVerify: true}),
		pgdriver.WithUser(user),
		pgdriver.WithPassword(password),
		pgdriver.WithDatabase(dbname))

	// setting the connector
	pdb.dbname = dbname
	pdb.connector = pgconn
}
