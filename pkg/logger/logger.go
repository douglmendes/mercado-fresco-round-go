package logger

import (
	"context"
	"log"
	"time"
)

type logger struct {
	level     string
	timestamp string
	caller    string
	message   string
}

const TimeLayout = "yyyy-MM-dd HH:mm:ss"

func Error(ctx context.Context, caller, message string) {
	timeNow := time.Now()
	timeNow.Format(TimeLayout)
	timeString := timeNow.String()
	timeString = timeString[:len(timeString)-32]

	l := logger{
		level:     "error",
		timestamp: timeString,
		caller:    caller,
		message:   message,
	}
	log.Println(l.level, l.timestamp, l.caller, l.message)
	CreateLog(ctx, l.level, l.timestamp, l.caller, l.message)
}

func Info(ctx context.Context, caller, message string) {
	timeNow := time.Now()
	timeNow.Format(TimeLayout)
	timeString := timeNow.String()
	timeString = timeString[:len(timeString)-32]

	l := logger{
		level:     "info",
		timestamp: timeString,
		caller:    caller,
		message:   message,
	}
	log.Println(l.level, l.timestamp, l.caller, l.message)
	CreateLog(ctx, l.level, l.timestamp, l.caller, l.message)
}

func Debug(ctx context.Context, caller, message string) {
	timeNow := time.Now()
	timeNow.Format(TimeLayout)
	timeString := timeNow.String()
	timeString = timeString[:len(timeString)-32]

	l := logger{
		level:     "debug",
		timestamp: timeString,
		caller:    caller,
		message:   message,
	}
	log.Println(l.level, l.timestamp, l.caller, l.message)
	CreateLog(ctx, l.level, l.timestamp, l.caller, l.message)
}
