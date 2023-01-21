package handler

import (
	"errors"
	"fmt"
	"github.com/rs/zerolog/log"
)

type Handler struct {
}

func HandleError(format string, args ...interface{}) error {
	msg := fmt.Sprintf(format, args...)
	log.Error().Msg(msg)
	return errors.New(msg)
}
