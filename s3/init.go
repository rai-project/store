package s3

import (
	"github.com/sirupsen/logrus"

	"github.com/rai-project/config"
	logger "github.com/rai-project/logger"
)

type logwrapper struct {
	*logrus.Entry
}

var (
	log *logwrapper
)

func (l *logwrapper) Log(args ...interface{}) {
	log.Debug(args...)
}

func init() {
	config.AfterInit(func() {
		log = &logwrapper{
			Entry: logger.New().WithField("pkg", "store/s3"),
		}
	})
}
