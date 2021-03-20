package stores

import "partyhann/backend/models"

//PartyStore ...
type PartyStore struct {
}

func NewPartyStore() *PartyStore {
	return &PartyStore{}
}

func (ps *PartyStore) Find(id int64) (models.Party, error) {
	return models.Party{}, nil
}
func (ps *PartyStore) All(offset int64, limit int64) ([]models.Party, error) {
	return []models.Party{}, nil
}
func (ps *PartyStore) Create(party models.Party) (int64, error) {
	return 0, nil
}
func (ps *PartyStore) Update(id int64, party models.Party) error {
	return nil
}
func (ps *PartyStore) Delete(id int64) error {
	return nil
}
