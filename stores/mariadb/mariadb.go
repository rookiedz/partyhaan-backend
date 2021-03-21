package mariadb

import (
	"database/sql"
	"fmt"
	"strings"
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
func GetCurrentDatetimeString() string {
	datetime := time.Now()
	return datetime.Format("2006-01-02 15:04:05")
}
func ArrayInt64ToString(arr []int64, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(arr), " ", delim, -1), "[]")
}
