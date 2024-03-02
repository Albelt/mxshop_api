package initial

import (
	"go.uber.org/zap"
)

func InitLogger() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	zap.ReplaceGlobals(logger)
}
