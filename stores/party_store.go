package stores

import (
	"partyhaan/backend/models"
	"partyhaan/backend/stores/mariadb"
)

//PartyStore ...
type PartyStore struct {
}

func NewPartyStore() *PartyStore {
	return &PartyStore{}
}

func (ps *PartyStore) Find(id int64) (models.Party, error) {
	mdbParty := mariadb.NewParty()
	return mdbParty.Find(id)
}
func (ps *PartyStore) FindByUser(id int64, offset int64, limit int64) ([]models.Party, error) {
	mdbParty := mariadb.NewParty()
	return mdbParty.FindByUser(id, offset, limit)
}
func (ps *PartyStore) All(offset int64, limit int64) ([]models.Party, error) {
	mdbParty := mariadb.NewParty()
	return mdbParty.All(offset, limit)
}
func (ps *PartyStore) Create(party models.Party) (int64, error) {
	mdbParty := mariadb.NewParty()
	return mdbParty.Create(party)
}
func (ps *PartyStore) Update(id int64, party models.Party) error {
	mdbParty := mariadb.NewParty()
	return mdbParty.Update(id, party)
}
func (ps *PartyStore) Delete(id int64) error {
	mdbParty := mariadb.NewParty()
	return mdbParty.Delete(id)
}
