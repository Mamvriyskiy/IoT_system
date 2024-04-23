package PPO

import (
	"go.uber.org/zap"
	//"time"
)

func Log(level, nameFunc, event string, err error, additionalParams ...interface{}) {
    logger, errZap := zap.NewDevelopment() 
    if errZap != nil {
        panic(errZap) // Не удалось создать логгер
    }

	switch level {
	case "Info":
	case "Error":
        logger.Error(
            err.Error(),
            zap.String("event", event),
            zap.String("func", nameFunc),
            zap.Any("param", additionalParams), 
		)
	case "Warning":
		logger.Warn(
            "event: " + event,
			zap.String("func", nameFunc),
			zap.Any("param", additionalParams),
		)
	}
}
