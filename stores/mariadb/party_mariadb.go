package mariadb

import (
	"context"
	"database/sql"
	"fmt"
	"partyhaan/backend/models"
	"strings"
	"time"
)

//PartyMariaDB ...
type PartyMariaDB struct {
	TableName    string
	Columns      []string
	InsertColumn string
	QueryColumn  string
}

func NewParty() *PartyMariaDB {
	columns := []string{
		"party_id",
		"party_conver",
		"party_name",
		"party_limit",
		"party_join",
		"party_owner",
	}
	return &PartyMariaDB{
		TableName:    "parties",
		Columns:      columns,
		InsertColumn: strings.Join(columns[1:], ","),
		QueryColumn:  strings.Join(columns[:], ","),
	}
}

func (pm *PartyMariaDB) Find(id int64) (models.Party, error) {
	var mParty models.Party
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	stmt, err := dbQuery.PrepareContext(ctx, fmt.Sprintf(`SELECT %s FROM %s WHERE party_id = ?`, pm.QueryColumn, pm.TableName))
	if err != nil {
		return mParty, err
	}
	defer stmt.Close()
	if err := stmt.QueryRowContext(ctx, id).Scan(&mParty.Cover, &mParty.Name, &mParty.Limit, &mParty.Join, &mParty.Owner); err != nil {
		if err == sql.ErrNoRows {
			return mParty, nil
		}
		return mParty, err
	}
	return mParty, nil
}
func (pm *PartyMariaDB) FindByUser(id int64, offset int64, limit int64) ([]models.Party, error) {
	var mParties []models.Party
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	stmt, err := dbQuery.PrepareContext(ctx, fmt.Sprintf(`SELECT %s FROM %s WHERE party_id IN (SELECT party_id FROM joins WHERE user_id = ?) ORDER BY party_id DESC LIMIT ?,?`, pm.QueryColumn, pm.TableName))
	if err != nil {
		return mParties, err
	}
	defer stmt.Close()
	rows, err := stmt.QueryContext(ctx, id, offset-1, limit)
	if err != nil {
		return mParties, err
	}
	defer rows.Close()
	for rows.Next() {
		var mParty models.Party
		if err := rows.Scan(&mParty.Cover, &mParty.Name, &mParty.Limit, &mParty.Join, &mParty.Owner); err != nil {
			return mParties, err
		}
		mParties = append(mParties, mParty)
	}
	if err = rows.Close(); err != nil {
		return mParties, err
	}
	if err = rows.Err(); err != nil {
		return mParties, err
	}
	return mParties, nil
}
func (pm *PartyMariaDB) All(offset int64, limit int64) ([]models.Party, error) {
	var mParties []models.Party
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	stmt, err := dbQuery.PrepareContext(ctx, fmt.Sprintf(`SELECT %s FROM %s ORDER BY party_id DESC LIMIT ?,?`, pm.QueryColumn, pm.TableName))
	if err != nil {
		return mParties, err
	}
	defer stmt.Close()
	rows, err := stmt.QueryContext(ctx, offset-1, limit)
	if err != nil {
		return mParties, err
	}
	defer rows.Close()
	for rows.Next() {
		var mParty models.Party
		if err := rows.Scan(&mParty.Cover, &mParty.Name, &mParty.Limit, &mParty.Join, &mParty.Owner); err != nil {
			return mParties, err
		}
		mParties = append(mParties, mParty)
	}
	if err = rows.Close(); err != nil {
		return mParties, err
	}
	if err = rows.Err(); err != nil {
		return mParties, err
	}
	return mParties, nil
}
func (pm *PartyMariaDB) Create(party models.Party) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	stmt, err := dbExec.PrepareContext(ctx, fmt.Sprintf(`INSERT INTO %s (%s, party_created, party_updated)VALUES(?,?,?,?,?,?,?)`, pm.InsertColumn, pm.TableName))
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	cds := GetCurrentDatetimeString()
	res, err := stmt.ExecContext(ctx, party.Cover, party.Name, party.Limit, party.Join, party.Owner, cds, cds)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}
func (pm *PartyMariaDB) Update(id int64, party models.Party) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	stmt, err := dbExec.PrepareContext(ctx, fmt.Sprintf(`UPDATE %s SET party_cover = ?, party_name = ?, party_limit = ?, party_join = ?, party_owner = ? WHERE party_id = ?`, pm.TableName))
	if err != nil {
		return err
	}
	defer stmt.Close()
	if _, err := stmt.ExecContext(ctx, party.Cover, party.Name, party.Limit, party.Join, party.Owner, id); err != nil {
		return err
	}
	return nil
}
func (pm *PartyMariaDB) Delete(id int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	stmt, err := dbExec.PrepareContext(ctx, fmt.Sprintf(`DELETE FROM %s WHERE party_id = ?`, pm.TableName))
	if err != nil {
		return err
	}
	defer stmt.Close()
	if _, err := stmt.ExecContext(ctx, id); err != nil {
		return err
	}
	return nil
}
