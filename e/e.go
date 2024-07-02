package e

import (
	log "github.com/sirupsen/logrus"
)

func Fatalf(format string, err error, args ...interface{}) {
	switch {
	case err != nil:
		log.Fatalf(format+" ACTION: Exit(1). ", err, args)
	}
}
