package datasource

import (
	"fmt"
	"github.com/go-sql-driver/mysql"
	"testing"
	"time"
)

func TestDSN(t *testing.T)  {
	cfg := mysql.Config{
		User:                 "root",
		Passwd:               "root",
		Addr:                 fmt.Sprintf("%s:%d", "localhost", 3306),
		DBName:               "test",
		Collation:            "utf8mb4_general_ci",
		ParseTime:            true,
		Loc:                  time.Local,
		Net:                  "tcp",
		AllowNativePasswords: true,
	}
	t.Log(cfg.FormatDSN())
}