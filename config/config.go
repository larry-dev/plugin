package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/larry-dev/plugins/eglog"
	"github.com/spf13/viper"
	"strings"
)

type Config struct {
	Name      string
	EnvPrefix string
	Type      string
	Watch     bool
}

// DefaultConfig 默认配置
func DefaultConfig() *Config {
	return &Config{
		EnvPrefix: "EG",
		Type:      "yaml",
		Watch:     true,
	}
}
func Init(c *Config) error {

	// 初始化配置文件
	if err := c.initConfig(); err != nil {
		return err
	}

	// 监控配置文件变化并热加载程序
	if c.Watch {
		c.watchConfig()
	}
	return nil
}

func (c *Config) initConfig() error {
	if c.Name != "" {
		viper.SetConfigFile(c.Name) // 如果指定了配置文件，则解析指定的配置文件
	} else {
		viper.AddConfigPath("conf") // 如果没有指定配置文件，则解析默认的配置文件
		viper.SetConfigName("config")
	}
	viper.SetConfigType(c.Type)     // 设置配置文件格式为YAML
	viper.AutomaticEnv()            // 读取匹配的环境变量
	viper.SetEnvPrefix(c.EnvPrefix) // 读取环境变量的前缀为EG
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	if err := viper.ReadInConfig(); err != nil { // viper解析配置文件
		return err
	}

	return nil
}

// 监控配置文件变化并热加载程序
func (c *Config) watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		eglog.Debug().Msgf("Config file changed: %s", e.Name)
	})
}
