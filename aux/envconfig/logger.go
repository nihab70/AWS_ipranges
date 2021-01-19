package envconfig

import (
	"os"

	log "github.com/sirupsen/logrus"
)

const envLogLevel = "LOG_LEVEL"
const defaultLogLevel = log.InfoLevel

// InitLogging determins the loglevel from $LOG_LEVEL or set default
func InitLogging() {
	level := defaultLogLevel

	// setting default values
	log.SetFormatter(&log.TextFormatter{})
	log.SetLevel(level)
	log.Infof("Logging is on default log level %v", level)

	// parsing LOG_LEVEL os environment variables
	levelString, exists := os.LookupEnv(envLogLevel)
	if exists {
		log.Infof("Loglevel read from OS environment %v", envLogLevel)
		level, err := log.ParseLevel(levelString)
		if err == nil {
			log.Infof("Logging is on log level %v", level)
			log.SetLevel(level)
		}
	}
}
