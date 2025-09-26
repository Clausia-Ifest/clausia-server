package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func NewZeroLog(appEnv string) error {
	timezone, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		return err
	}

	if appEnv == "prod" || appEnv == "production" {
		writer, err := rotatelogs.New(
			"./logs/%Y-%m-%d.log",
			rotatelogs.WithLinkName("./logs/app.log"),
			rotatelogs.WithRotationTime(24*time.Hour),
		)

		if err != nil {
			return err
		}

		zerolog.TimeFieldFormat = time.RFC3339
		log.Logger = zerolog.New(writer).
			With().
			Timestamp().
			Logger()
	} else {
		consoleWriter := zerolog.ConsoleWriter{
			Out:          os.Stdout,
			TimeFormat:   time.RFC1123,
			PartsOrder:   []string{"level", "message", "time"},
			FieldsOrder:  []string{"trace_id", "ip", "operation", "duration", "log_info"},
			TimeLocation: timezone,
			FormatCaller: func(i interface{}) string {
				var c string
				if cc, ok := i.(string); ok {
					c = cc
				}
				if len(c) > 0 {
					if cwd, err := os.Getwd(); err == nil {
						if rel, err := filepath.Rel(cwd, c); err == nil {
							c = rel
						}
					}
				}
				if c != "" {
					return fmt.Sprintf("[%v]", c)
				}
				return c
			},
		}
		log.Logger = zerolog.New(consoleWriter).With().Timestamp().Logger()
	}

	return nil
}
