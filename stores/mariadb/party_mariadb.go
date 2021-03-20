package mariadb

import "partyhann/backend/models"

//PartyMariaDB ...
type PartyMariaDB struct{}

func (pm *PartyMariaDB) Find(id int64) (models.Party, error) {
	return models.Party{}, nil
}
func (pm *PartyMariaDB) All(offset int64, limit int64) ([]models.Party, error) {
	return []models.Party{}, nil
}
func (pm *PartyMariaDB) Create(party models.Party) (int64, error) {
	return 0, nil
}
func (pm *PartyMariaDB) Update(id int64, party models.Party) error {
	return nil
}
func (pm *PartyMariaDB) Delete(id int64) error {
	return nil
}
