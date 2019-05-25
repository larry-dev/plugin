package eglog

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"testing"
)

func TestInitLog(t *testing.T) {
	viper.SetConfigType("yaml")
	viper.SetConfigName("test")
	viper.AddConfigPath(".")
	if err:=viper.ReadInConfig();err!=nil{
		t.Error(err)
	}
	if err:=InitLog();err!=nil{
		t.Error(err)
	}
	log.Info().Msg("测试")
}