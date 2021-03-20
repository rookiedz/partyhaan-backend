package mariadb

import "partyhann/backend/models"

//AuthenMariaDB ...
type AuthenMariaDB struct{}

func (am *AuthenMariaDB) FindByEmail(email string) (models.Authen, error) {
	return models.Authen{}, nil
}
func (am *AuthenMariaDB) Create(models.Authen) (int64, error) {
	return 0, nil
}
func (am *AuthenMariaDB) Update(id int64, authen models.Authen) error {
	return nil
}
func (am *AuthenMariaDB) Delete(id int64) error {
	return nil
}
