package main

import (
	"fmt"
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

	// cleaning the DB
	err = helper.CleanDB()
	if err != nil {
		msg := fmt.Sprintf("error cleaning: %v", err)
		log.Error().Msg(msg)
	}

	// creating the table
	err = helper.CreateSchemas()
	if err != nil {
		msg := fmt.Sprintf("error creating: %v", err)
		log.Error().Msg(msg)
	}

	// dumping the players
	err = helper.DumpPlayersToDB()
	if err != nil {
		msg := fmt.Sprintf("error dumping: %v", err)
		log.Error().Msg(msg)
	}

	api.StartRouter(&db)
}
