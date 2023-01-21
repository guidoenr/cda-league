package main

import (
	"github.com/guidoenr/fulbo/api"
	"github.com/guidoenr/fulbo/model/psdb"
	"github.com/rs/zerolog/log"
)

func main() {
	// initializing the database
	var db psdb.PostgreDB
	err := db.InitDB()
	if err != nil {
		log.Error().Msgf("initializing db: %v", err)
	}
	defer db.CloseDB()

	// creating the helper
	var helper api.Helper
	helper.Init(&db)

	helper.InitializeDatabase()

	var Router api.Router
	Router.Init(&db)

	Router.StartRouter()
}
