package datasource

import (
	"fmt"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/spf13/viper"
	"time"
)

func newMysql() (*xorm.Engine, error) {
	cfg := mysql.Config{
		User:                 viper.GetString("mysql.user"),
		Passwd:               viper.GetString("mysql.passwd"),
		Addr:                 fmt.Sprintf("%s:%d", viper.GetString("mysql.host"), viper.GetInt("mysql.port")),
		DBName:               viper.GetString("mysql.db"),
		Collation:            "utf8mb4_general_ci",
		ParseTime:            true,
		Loc:                  time.Local,
		Net:                  "tcp",
		AllowNativePasswords: true,
	}
	engine, err := xorm.NewEngine("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, err
	}
	if viper.GetBool("mysql.log") {
		engine.ShowSQL(true)
	}
	engine.SetMaxIdleConns(viper.GetInt("mysql.min"))
	engine.SetMaxOpenConns(viper.GetInt("mysql.max"))
	if err = engine.Ping(); err != nil {
		return nil, err
	}
	return engine, nil
}
