package eglog

import (
	"encoding/json"
	"github.com/arthurkiller/rollingwriter"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"os"
	"sync"
)

var (
	once   sync.Once
	Logger zerolog.Logger
)

func InitLog() {
	once.Do(func() {
		if viper.GetBool("debug") {
			zerolog.SetGlobalLevel(zerolog.DebugLevel)
		} else {
			zerolog.SetGlobalLevel(zerolog.InfoLevel)
		}
		zerolog.TimeFieldFormat = "2006-01-02 15:04:05"
		var logAdapter = viper.GetString("log.adapter")
		switch logAdapter {
		case "color":
			Logger = zerolog.New(zerolog.ConsoleWriter{
				Out:        os.Stdout,
				NoColor:    false,
				TimeFormat: "2006-01-02 15-04-05",
			})
		case "":
			Logger = zerolog.New(os.Stdout)
		case "file":
			config := rollingwriter.NewDefaultConfig()
			form := viper.GetString("log.format")
			if form == "" {
				form = "{}"
			}
			if err := json.Unmarshal([]byte(form), &config); err != nil {
				log.Error().Err(err)
			}
			writer, err := rollingwriter.NewWriterFromConfig(&config)
			if err != nil {
				log.Error().Err(err)
			}
			Logger = zerolog.New(writer)
		}
		initLogger()
	})
}

/**
绑定log属性
*/
func initLogger() {
	Logger = Logger.With().Timestamp().Logger()
}
func WithCaller()  {
	Logger=Logger.With().Caller().Logger()
}
func Error() *zerolog.Event {
	return Logger.Error()
}
func Info() *zerolog.Event {
	return Logger.Info()
}
func Debug() *zerolog.Event {
	return Logger.Debug()
}
func Fatal() *zerolog.Event {
	return Logger.Fatal()
}
func Warn() *zerolog.Event {
	return Logger.Warn()
}
