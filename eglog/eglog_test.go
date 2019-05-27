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
	Logger.Print("测试")
}
