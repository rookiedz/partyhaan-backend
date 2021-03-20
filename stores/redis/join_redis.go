package redis

import "partyhann/backend/models"

//JoinRedis ...
type JoinRedis struct{}

func (jr *JoinRedis) FindByUser(id int64) ([]models.Party, error) {
	return []models.Party{}, nil
}
func (jr *JoinRedis) FindByParty(id int64) ([]models.User, error) {
	return []models.User{}, nil
}
func (jr *JoinRedis) Create(join models.Join) (int64, error) {
	return 0, nil
}
func (jr *JoinRedis) Delete(id int64) error {
	return nil
}
