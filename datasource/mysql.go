package datasource

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"time"
)

func newMysql() (*sql.DB, error) {
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
	engine, err := gorm.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, err
	}
	if viper.GetBool("mysql.log") {
		engine.LogMode(true)
	}
	engine.DB().SetMaxIdleConns(viper.GetInt("mysql.min"))
	engine.DB().SetMaxOpenConns(viper.GetInt("mysql.max"))
	if err = engine.DB().Ping(); err != nil {
		return nil, err
	}
	return engine.DB(), nil
}
