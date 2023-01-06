package psdb

import (
	"crypto/tls"
	"database/sql"
	"fmt"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"os"
)

type PostgreDB struct {
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
	db := bun.NewDB(sqldb, pgdialect.New())

	// pinging to DB
	err := db.Ping()
	if err != nil {
		return fmt.Errorf("pinging to db: %v", err)
	}

	fmt.Printf("connected to db '%s' succesfully", pdb.dbname)
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
