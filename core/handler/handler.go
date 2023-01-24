package h

import (
	"fmt"
	"github.com/fatih/color"
	"log"
	"runtime"
)

// HandleError handles the different errors code to perform the required actions or just
// print the errors in a different way
func HandleError(err Error) error {
	switch err.Code {
	case DbError:
		log.Fatal(color.RedString("[ERROR-DB]: %v", err.Error()))
	case RouterError:
		log.Fatal(color.YellowString("[ERROR-ROUTER]: %v", err.Error()))
	case FileError:
		log.Fatal(color.HiRedString("[ERROR-FILE]: %v", err.Error()))
	case PlayerError:
	case MatchError:
	case ApiError:

	default:
		// handle unknown error
	}
	return nil
}

// LogInfo just prints the info
func LogInfo(msg string, args ...interface{}) {
	log.Printf(color.BlueString("[INFO]:"+msg), args...)
}

func LogOk(msg string, args ...interface{}) {
	log.Printf(color.HiGreenString("[OK]:"+msg), args...)
}

// -------------------------------- ERRORS
// Errors const
const (
	// DbError = iota makes the error to be incremental, DbError = 0, AuthError=1, etc
	DbError = iota
	// RouterError = 1
	RouterError
	// ApiError = 2
	ApiError
	// BackupError
	BackupError
	// FileError = 3
	FileError
	// PlayerError = 4
	PlayerError
	// MatchError = 5
	MatchError
)

type Error struct {
	Code    int
	Message string
	Stack   []uintptr
}

func (e *Error) Error() string {
	return e.Message
}

func New(code int, message string) Error {
	var stack []uintptr
	n := runtime.Callers(2, stack[:])
	return Error{
		Code:    code,
		Message: message,
		Stack:   stack[:n],
	}
}

func Newf(code int, format string, args ...interface{}) Error {
	return New(code, fmt.Sprintf(format, args...))
}
