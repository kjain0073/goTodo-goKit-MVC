package adapters

import (
	"os"

	"github.com/go-kit/log"
)

func InitLogger() log.Logger {
	var logger log.Logger

	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.NewSyncLogger(logger)
	logger = log.With(logger,
		"service", "tasks",
		"time:", log.DefaultTimestampUTC,
		"caller", log.DefaultCaller,
	)
	return logger
}
