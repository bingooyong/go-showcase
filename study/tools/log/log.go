package log

import (
	"study/tools/config"
	"os"
	"github.com/sirupsen/logrus"
)

func Config(c *config.Config) {
	logFile := c.LogPath
	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY, 0777)
	if err == nil {
		logrus.StandardLogger().Out = file
	} else {
		logrus.Info("Failed to log to file, using default stderr")
	}
}
