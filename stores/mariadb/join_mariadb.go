package mariadb

import (
	"context"
	"fmt"
	"partyhaan/backend/models"
	"strings"
	"time"
)

//JoinMariaDB ...
type JoinMariaDB struct {
	TableName    string
	Columns      []string
	InsertColumn string
	QueryColumn  string
}

func NewJoin() *JoinMariaDB {
	colums := []string{
		"join_id",
		"party_id",
		"user_id",
	}
	return &JoinMariaDB{
		TableName:    "joins",
		Columns:      colums,
		InsertColumn: strings.Join(colums[1:], ","),
		QueryColumn:  strings.Join(colums[:], ","),
	}
}

func (jm *JoinMariaDB) Create(join models.Join) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	stmt, err := dbExec.PrepareContext(ctx, fmt.Sprintf(`INSERT INTO %s (%s,join_created, join_updated)VALUES(?,?,?,?)`, jm.TableName, jm.InsertColumn))
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	cds := GetCurrentDatetimeString()
	res, err := stmt.ExecContext(ctx, join.PartyID, join.UserID, cds)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}
func (jm *JoinMariaDB) Delete(id int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	stmt, err := dbExec.PrepareContext(ctx, fmt.Sprintf(`DELETE FROM %s WHERE join_id = ?`, jm.TableName))
	if err != nil {
		return err
	}
	defer stmt.Close()
	if _, err := stmt.ExecContext(ctx, id); err != nil {
		return err
	}
	return nil
}
