package config

import (
	"encoding/json"
	"github.com/rs/zerolog/log"
	"os"
)

type Config struct {
	TeamNames []string `json:"TeamNames"`
}

func (c *Config) Init() {
	bytes, err := os.ReadFile("config/config.json")
	if err != nil {
		log.Error().Msgf("reading file: %v", err)
	}

	// TODO, not breaking everytihng just for this, then
	// i have to refactor this and see..
	err = json.Unmarshal(bytes, &c)
	if err != nil {
		log.Error().Msgf("unmarshalling file: %v", err)
	}
}
