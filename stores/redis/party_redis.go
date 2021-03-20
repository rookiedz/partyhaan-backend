package redis

import "partyhann/backend/models"

//PartyRedis ...
type PartyRedis struct{}

func (pr *PartyRedis) Find(id int64) (models.Party, error) {
	return models.Party{}, nil
}
func (pr *PartyRedis) All(offset int64, limit int64) ([]models.Party, error) {
	return []models.Party{}, nil
}
func (pr *PartyRedis) Create(party models.Party) (int64, error) {
	return 0, nil
}
func (pr *PartyRedis) Update(id int64, party models.Party) error {
	return nil
}
func (pr *PartyRedis) Delete(id int64) error {
	return nil
}
