package mariadb

import (
	"context"
	"database/sql"
	"fmt"
	"partyhaan/backend/models"
	"strings"
	"time"
)

//AuthenMariaDB ...
type AuthenMariaDB struct {
	TableName    string
	Columns      []string
	InsertColumn string
	QueryColumn  string
}

func NewAuthen() *AuthenMariaDB {
	columns := []string{
		"authen_id",
		"authen_email",
		"authen_password",
		"authen_salt",
		"authen_hash",
		"user_id",
	}
	return &AuthenMariaDB{
		TableName:    "authens",
		Columns:      columns,
		InsertColumn: strings.Join(columns[1:], ","),
		QueryColumn:  strings.Join(columns[:], ","),
	}
}

func (am *AuthenMariaDB) FindByEmail(email string) (models.Authen, error) {
	var mAuthen models.Authen
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	stmt, err := dbQuery.PrepareContext(ctx, fmt.Sprintf(`SELECT %s FROM %s WHERE authen_email = ?`, am.QueryColumn, am.TableName))
	if err != nil {
		return mAuthen, err
	}
	defer stmt.Close()
	if err := stmt.QueryRowContext(ctx, email).Scan(&mAuthen.ID, &mAuthen.Email, &mAuthen.Password, &mAuthen.Salt, &mAuthen.Hash, &mAuthen.UserID); err != nil {
		if err == sql.ErrNoRows {
			return mAuthen, nil
		}
		return mAuthen, err
	}
	return mAuthen, nil
}
func (am *AuthenMariaDB) Create(authen models.Authen) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	stmt, err := dbExec.PrepareContext(ctx, fmt.Sprintf(`INSERT INTO %s (%s, authen_created, authen_updated)VALUES(?,?,?,?,?,?,?)`, am.TableName, am.InsertColumn))
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	cds := GetCurrentDatetimeString()
	res, err := stmt.ExecContext(ctx, authen.Email, authen.Password, authen.Salt, authen.Hash, authen.UserID, cds, cds)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}
func (am *AuthenMariaDB) Update(id int64, authen models.Authen) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt, err := dbExec.PrepareContext(ctx, fmt.Sprintf(`UPDATE %s SET authen_email = ?, authen_password = ?, authen_salt = ?, authen_hash = ?, user_id = ?, authen_updated = ? WHERE authen_id = ?`, am.TableName))
	if err != nil {
		return err
	}
	defer stmt.Close()
	cds := GetCurrentDatetimeString()
	if _, err := stmt.ExecContext(ctx, authen.Email, authen.Password, authen.Salt, authen.Hash, authen.UserID, cds, id); err != nil {
		return err
	}
	return nil
}
func (am *AuthenMariaDB) Delete(id int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt, err := dbExec.PrepareContext(ctx, fmt.Sprintf(`DELETE FROM %s WHERE authen_id = ?`, am.TableName))
	if err != nil {
		return err
	}
	defer stmt.Close()
	if _, err := stmt.ExecContext(ctx, id); err != nil {
		return err
	}
	return nil
}
