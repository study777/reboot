package mysql

import (
	"fmt"
	"github.com/jmoiron/sqlx"

	"reboot/pkg/dao"
)

type Options struct {
	DbConnStr string
	DbName    string
}

type mysql struct {
	opt *Options
	db  *sqlx.DB
}

func New(opt *Options) dao.Storage {
	db, err := sqlx.Open("mysql", opt.DbConnStr)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		return nil
	}

	s := mysql{
		opt: opt,
		db:  db,
	}
	return &s
}
