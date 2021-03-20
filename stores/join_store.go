package stores

import "partyhann/backend/models"

//JoinStore ...
type JoinStore struct{}

func (js *JoinStore) FindByUser(id int64) ([]models.Party, error) {
	return []models.Party{}, nil
}
func (js *JoinStore) FindByParty(id int64) ([]models.User, error) {
	return []models.User{}, nil
}
func (js *JoinStore) Create(join models.Join) (int64, error) {
	return 0, nil
}
func (js *JoinStore) Delete(id int64) error {
	return nil
}
