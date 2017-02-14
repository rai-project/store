package s3

import (
	"github.com/Sirupsen/logrus"

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
	config.OnInit(func() {
		log = &logwrapper{
			Entry: logger.New().WithField("pkg", "s3-store"),
		}
	})
}
