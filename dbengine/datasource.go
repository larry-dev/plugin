package datasource

import (
	"fmt"
	"github.com/spf13/viper"
)

// 获取mysql DSN
func mysqlDSN() string {

	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?collation=utf8mb4_general_ci&loc=Local&parseTime=true&maxAllowedPacket=0",
		viper.GetString("mysql.user"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.host"),
		viper.GetInt("mysql.port"),
		viper.GetString("mysql.db_name"))
}

// 获取postgres DSN
func postgresDSN() string {
	return fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s",
		viper.GetString("postgres.host"),
		viper.GetInt("postgres.port"),
		viper.GetString("postgres.user"),
		viper.GetString("postgres.db_name"),
		viper.GetString("postgres.password"))
}

// 获取sqlite DSN
func sqliteDSN() string{
	if dsn := viper.GetString("sqlite.path"); dsn != "" {
		return dsn
	} else {
		return "./egcode.db"
	}
}
