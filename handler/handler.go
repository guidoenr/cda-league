package handler

import (
	"errors"
	"fmt"
	"github.com/rs/zerolog/log"
)

// Handler handles all the errors
type Handler struct {
}

func (h Handler) HandleError(format string, args ...interface{}) error {
	msg := fmt.Sprintf(format, args...)
	log.Error().Msg(msg)
	return errors.New(msg)
}
