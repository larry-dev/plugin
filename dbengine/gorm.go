package datasource

import (
	"database/sql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"time"
)

func NewDBEngine() (*sql.DB, error) {
	dbType := viper.GetString("gorm.dbtype")
	if dbType == "" {
		dbType = "sqlite"
	}

	dsn := ""
	switch dbType {
	case "mysql":
		dsn = mysqlDSN()
	case "postgres":
		dsn = postgresDSN()
	case "sqlite":
		dsn = sqliteDSN()
	default:
		return nil, errors.New("only suppert mysql postgres sqlite db")
	}
	engine, err := gorm.Open(dbType, dsn)
	if err != nil {
		return nil, err
	}
	engine.LogMode(viper.GetBool("gorm.debug"))
	engine.DB().SetMaxIdleConns(viper.GetInt("gorm.min"))
	engine.DB().SetMaxOpenConns(viper.GetInt("gorm.max"))
	engine.DB().SetConnMaxLifetime(time.Duration(viper.GetInt("gorm.life_time")) * time.Second)
	if err = engine.DB().Ping(); err != nil {
		return nil, err
	}
	return engine.DB(), nil
}
