package stores

import "partyhann/backend/models"

//UserStore ...
type UserStore struct{}

func (us *UserStore) Find(id int64) (models.User, error) {
	return models.User{}, nil
}
func (us *UserStore) Create(user models.User) (int64, error) {
	return 0, nil
}
func (us *UserStore) Update(id int64, user models.User) error {
	return nil
}
func (us *UserStore) Delete(id int64) error {
	return nil
}
