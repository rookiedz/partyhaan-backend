package mariadb

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var dbQuery *sql.DB
var dbExec *sql.DB

func connect(dsn string) *sql.DB {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(time.Second * 30)
	db.SetMaxIdleConns(50)
	db.SetMaxOpenConns(50)
	return db
}
func QueryConnect(dsn string) {
	dbQuery = connect(dsn)
}
func ExecConnect(dsn string) {
	dbExec = connect(dsn)
}
