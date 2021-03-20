package redis

import "partyhann/backend/models"

//AuthenRedis ...
type AuthenRedis struct{}

func (ar *AuthenRedis) FindByEmail(email string) (models.Authen, error) {
	return models.Authen{}, nil
}
func (ar *AuthenRedis) Create(authen models.Authen) (int64, error) {
	return 0, nil
}
func (ar *AuthenRedis) Update(id int64, authen models.Authen) error {
	return nil
}
func (ar *AuthenRedis) Delete(id int64) error {
	return nil
}
