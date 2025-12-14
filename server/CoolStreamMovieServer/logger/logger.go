// Package logger provides a singalten logger
package logger

import (
	"os"
	"path/filepath"
	"strconv"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var zLogger *zerolog.Logger

func GetLogger() *zerolog.Logger {
	if zLogger == nil {
		zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
		zerolog.CallerMarshalFunc = func(pc uintptr, file string, line int) string {
			return filepath.Base(file) + ":" + strconv.Itoa(line)
		}
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
		log.Logger = log.With().Caller().Logger()
		zLogger = &log.Logger
	}
	return zLogger
}
