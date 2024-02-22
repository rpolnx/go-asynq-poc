package configs

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func InitLogger() {
	if GlobalAppConfig.Profile == LocalProfile {
		log.SetFormatter(&log.TextFormatter{ForceColors: true, FullTimestamp: true})
		log.SetOutput(os.Stdout)
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetFormatter(&log.JSONFormatter{})
		log.SetOutput(os.Stdout)
		log.SetLevel(log.WarnLevel)
	}
}
