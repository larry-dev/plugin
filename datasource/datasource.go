package datasource

import "github.com/go-xorm/xorm"

type DataSource struct {
	DB *xorm.Engine
}

func NewDataSource() (ds *DataSource, err error) {
	db, err := newMysql()
	if err != nil {
		return
	}
	ds = &DataSource{
		DB: db,
	}
	return
}
