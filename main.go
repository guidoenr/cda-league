package main

import (
	"github.com/guidoenr/cda-league/api"
	"github.com/guidoenr/cda-league/model/psdb"
	"github.com/rs/zerolog/log"
	"os"
)

func main() {
	// creating the database
	var db psdb.PostgreDB
	err := db.InitDB()
	if err != nil {
		log.Error().Msgf("initializing db: %v", err)
		os.Exit(0)
	}
	defer db.CloseDB()

	// creating the helper
	var helper api.Helper
	helper.Init(&db)

	// initializing the database
	helper.InitializeDatabase(true)

	// creating the router and their routes
	var Router api.Router
	Router.Init(&db)

	Router.StartRouter()
}
