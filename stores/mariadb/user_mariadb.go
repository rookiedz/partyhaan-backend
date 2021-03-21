package mariadb

import (
	"context"
	"database/sql"
	"fmt"
	"partyhaan/backend/models"
	"strings"
	"time"
)

//UserMariaDB ...
type UserMariaDB struct {
	TableName    string
	Columns      []string
	InsertColumn string
	QueryColumn  string
}

func NewUser() *UserMariaDB {
	columns := []string{"user_id", "user_fullname"}
	return &UserMariaDB{
		TableName:    "users",
		Columns:      columns,
		InsertColumn: strings.Join(columns[1:], ","),
		QueryColumn:  strings.Join(columns[:], ","),
	}
}

func (um *UserMariaDB) Find(id int64) (models.User, error) {
	var mUser models.User
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	stmt, err := dbQuery.PrepareContext(ctx, fmt.Sprintf(`SELECT %s FROM %s WHERE user_id = ?`, um.QueryColumn, um.TableName))
	if err != nil {
		return mUser, err
	}
	defer stmt.Close()
	if err := stmt.QueryRowContext(ctx, id).Scan(&mUser.ID); err != nil {
		if err == sql.ErrNoRows {
			return mUser, nil
		}
		return mUser, err
	}

	return mUser, nil
}
func (um *UserMariaDB) FindByParty(id int64, offset int64, limit int64) ([]models.User, error) {
	var mUsers []models.User
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	stmt, err := dbQuery.PrepareContext(ctx, fmt.Sprintf(`SELECT %s FROM %s WHERE user_id IN(SELECT user_id FROM joins WHERE party_id = ?) ORDER BY user_id LIMIT ?,?`, um.QueryColumn, um.TableName))
	if err != nil {
		return mUsers, err
	}
	defer stmt.Close()
	rows, err := stmt.QueryContext(ctx, id, offset-1, limit)
	if err != nil {
		return mUsers, err
	}
	defer rows.Close()
	for rows.Next() {
		var mUser models.User
		if err := rows.Scan(&mUser.Fullname); err != nil {
			return mUsers, err
		}
		mUsers = append(mUsers, mUser)
	}
	if err = rows.Close(); err != nil {
		return mUsers, err
	}
	if err = rows.Err(); err != nil {
		return mUsers, err
	}
	return mUsers, nil
}
func (um *UserMariaDB) Create(user models.User) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	stmt, err := dbExec.PrepareContext(ctx, fmt.Sprintf(`INSERT INTO %s (%s, user_created, user_updated)VALUES(?,?,?)`, um.TableName, um.InsertColumn))
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	cds := GetCurrentDatetimeString()
	res, err := stmt.ExecContext(ctx, user.Fullname, cds, cds)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}
func (um *UserMariaDB) Update(id int64, user models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	stmt, err := dbExec.PrepareContext(ctx, fmt.Sprintf(`UPDATE %s SET user_fullname = ?, user_updated = ? WHERE user_id = ?`, um.TableName))
	if err != nil {
		return err
	}
	defer stmt.Close()
	cds := GetCurrentDatetimeString()
	if _, err := stmt.ExecContext(ctx, user.Fullname, cds, id); err != nil {
		return err
	}
	return nil
}
func (um *UserMariaDB) Delete(id int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	stmt, err := dbExec.PrepareContext(ctx, fmt.Sprintf(`DELETE FROM %s WHERE user_id = ?`, um.TableName))
	if err != nil {
		return err
	}
	defer stmt.Close()
	if _, err := stmt.ExecContext(ctx, id); err != nil {
		return err
	}
	return nil
}
