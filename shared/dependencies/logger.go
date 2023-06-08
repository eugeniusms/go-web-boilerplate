package depedencies

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

func NewLogger() (*logrus.Logger, error) {
	var log = logrus.New()

	file, err := os.OpenFile("./logger.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	log.SetFormatter(&logrus.JSONFormatter{})
	mw := io.MultiWriter(os.Stdout, file)
	log.SetOutput(mw)

	return log, nil
}