package logging

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
)

func SetupLogrus(lvl string) error {
	level, err := logrus.ParseLevel(lvl)
	if err != nil {
		return fmt.Errorf("fail to parse log level: %w", err)
	}

	logrus.SetLevel(level)
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.DateTime,
	})

	return nil
}
