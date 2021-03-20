package stores

import "partyhann/backend/models"

//AuthenStore ...
type AuthenStore struct {
}

func (as *AuthenStore) FindByEmail(email string) (models.Authen, error) {
	return models.Authen{}, nil
}
func (as *AuthenStore) Create(authen models.Authen) (int64, error) {
	return 0, nil
}
func (as *AuthenStore) Update(id int64, authen models.Authen) error {
	return nil
}
func (as *AuthenStore) Delete(id int64) error {
	return nil
}
