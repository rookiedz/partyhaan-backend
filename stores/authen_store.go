package stores

import (
	"partyhaan/backend/models"
	"partyhaan/backend/stores/mariadb"
)

//AuthenStore ...
type AuthenStore struct {
}

func NewAuthenStore() *AuthenStore {
	return &AuthenStore{}
}

func (as *AuthenStore) FindByEmail(email string) (models.Authen, error) {
	mdbAuthen := mariadb.NewAuthen()
	return mdbAuthen.FindByEmail(email)
}
func (as *AuthenStore) Create(authen models.Authen) (int64, error) {
	mdbAuthen := mariadb.NewAuthen()
	return mdbAuthen.Create(authen)
}
func (as *AuthenStore) Update(id int64, authen models.Authen) error {
	mdbAuthen := mariadb.NewAuthen()
	return mdbAuthen.Update(id, authen)
}
func (as *AuthenStore) Delete(id int64) error {
	mdbAuthen := mariadb.NewAuthen()
	return mdbAuthen.Delete(id)
}
