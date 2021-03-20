package mariadb

import "partyhann/backend/models"

//JoinMariaDB ...
type JoinMariaDB struct{}

func (jm *JoinMariaDB) FindByUser(id int64) ([]models.Party, error) {
	return []models.Party{}, nil
}
func (jm *JoinMariaDB) FindByParty(id int64) (models.User, error) {
	return models.User{}, nil
}
func (jm *JoinMariaDB) Create(join models.Join) (int64, error) {
	return 0, nil
}
func (jm *JoinMariaDB) Delete(id int64) error {
	return nil
}
