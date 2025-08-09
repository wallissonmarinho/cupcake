package util

import (
	"fmt"

	"github.com/go-kit/log"
)

func Error(logger log.Logger, msg string, err error) {
	if errLog := logger.Log(
		"level", "error",
		"msg", msg,
		"error", err,
	); errLog != nil {
		fmt.Println("logger error:", errLog)
	}
}

func Info(logger log.Logger, msg string) {
	if errLog := logger.Log(
		"level", "info",
		"msg", msg,
	); errLog != nil {
		fmt.Println("logger error:", errLog)
	}
}
