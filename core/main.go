package main

import (
	api2 "github.com/guidoenr/cda-league/core/api"
	"github.com/guidoenr/cda-league/core/model/psdb"
	"github.com/rs/zerolog/log"
	"os"
)

func main() {
	// creating the database
	var db psdb.PostgreDB
	err := db.InitDB()
	if err != nil {
		log.Error().Msgf("creating postgredb connector: %v", err)
		os.Exit(0)
	}
	defer db.CloseDB()

	// creating the helper
	var helper psdb.DbManager
	helper.Init(&db)

	// initializing the database
	err = helper.InitializeDatabase()
	if err != nil {
		log.Error().Msgf("initializing db: %v", err)
		os.Exit(0)
	}

	// creating the router and their routes
	var Router api2.Router
	Router.Init(&db)

	Router.StartRouter()
}
