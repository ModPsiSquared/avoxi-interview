package utility

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func InitLogging() (cleanup func() error, err error) {
	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return
	}
	cleanup = file.Close

	log.SetOutput(file)
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.WarnLevel)
	return
}
