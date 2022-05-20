package mysql

import (
	"context"

	xsql "bitbucket.org/pwq/tata/lib/db/mysql"
)

const (
	_where = " WHERE "
	_and   = " AND "
	_or    = " OR "
	_end   = ";"
)

// Dao dao
type Dao struct {
	DB *xsql.DB
}

// New init mysql DB
func New(dsn xsql.Config) (dao *Dao) {
	dao = &Dao{
		DB: xsql.NewMySQL(&dsn),
	}

	return
}

// Close close the resource.
func (dao *Dao) Close() {
	_ = dao.DB.Close()
}

// Ping dao ping
func (dao *Dao) Ping(ctx context.Context) error {
	return dao.DB.Ping(ctx)
}
