package redis

import "partyhann/backend/models"

//UserRedis ...
type UserRedis struct{}

func (ur *UserRedis) Find(id int64) (models.User, error) {
	return models.User{}, nil
}
func (ur *UserRedis) Create(user models.User) (int64, error) {
	return 0, nil
}
func (ur *UserRedis) Update(id int64, user models.User) error {
	return nil
}
func (ur *UserRedis) Delete(id int64) error {
	return nil
}
