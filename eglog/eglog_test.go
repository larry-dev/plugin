package eglog

import (
	"github.com/spf13/viper"
	"testing"
)

func TestInitLog(t *testing.T) {
	viper.SetConfigType("yaml")
	viper.SetConfigName("test")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		t.Error(err)
	}
	InitLog()
	WithCaller()
	Logger.Print("测试")
	//Info().Msg("测试")
	Error().Msg("测试测物")
}
