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
	// create the connector
	var connector *pgdriver.Connector

	// get the vars
	dbname := os.Getenv("DB_NAME")

	// note that the DATABASE_URL is the same thing that the DSN
	// something like postgres://postgres:@localhost:5432/test
	dbUrl := os.Getenv("DATABASE_URL")

	// if the DATABASE_URL is contains something, that means we are in the
	// cloud environment, so we will use the DSN option (database URL)
	if dbUrl != "" {
		connector = pgdriver.NewConnector(
			pgdriver.WithDSN(dbUrl),
		)
	} else {
		connector = pgdriver.NewConnector(
			pgdriver.WithNetwork(os.Getenv("DB_NETWORK")),
			pgdriver.WithAddr(os.Getenv("DB_ADDR")),
			pgdriver.WithInsecure(true),
			pgdriver.WithTLSConfig(&tls.Config{InsecureSkipVerify: true}),
			pgdriver.WithUser(os.Getenv("DB_USER")),
			pgdriver.WithPassword(os.Getenv("DB_PASSWORD")),
			pgdriver.WithDatabase(os.Getenv("DB_NAME")))
	}

	// setting the connector
	pdb.dbname = dbname
	pdb.connector = connector
}
