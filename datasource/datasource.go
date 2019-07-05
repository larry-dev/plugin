package datasource

import (
	"database/sql"
)

type DataSource struct {
	MySql *sql.DB
}

func NewDataSource() (ds *DataSource, err error) {
	db, err := newMysql()
	if err != nil {
		return
	}
	ds = &DataSource{
		MySql: db,
	}
	return
}
