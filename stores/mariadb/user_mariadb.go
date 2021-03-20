package mariadb

import "partyhann/backend/models"

//UserMariaDB ...
type UserMariaDB struct{}

func (um *UserMariaDB) Find(id int) (models.User, error) {
	return models.User{}, nil
}
func (um *UserMariaDB) Create(user models.User) (int64, error) {
	return 0, nil
}
func (um *UserMariaDB) Update(id int64, user models.User) error {
	return nil
}
func (um *UserMariaDB) Delete(id int64) error {
	return nil
}
