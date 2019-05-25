package eglog

import (
	"encoding/json"
	"github.com/arthurkiller/rollingwriter"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"os"
)

func InitLog() error {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	var logAdapter=viper.GetString("log.adapter")
	switch logAdapter {
	case "console","":
		log.Logger=log.Output(zerolog.ConsoleWriter{
			Out:     os.Stdin,
			NoColor: false,
			TimeFormat:"2006-01-02 15-04-05",
		})
	case "file":
		config:=rollingwriter.NewDefaultConfig()
		form:=viper.GetString("log.format")
		if form==""{
			form="{}"
		}
		if err := json.Unmarshal([]byte(form), &config);err!=nil{
			return err
		}
		writer,err:=rollingwriter.NewWriterFromConfig(&config)
		if err!=nil{
			return err
		}
		log.Logger=log.Output(writer)
	}
	return nil
}
