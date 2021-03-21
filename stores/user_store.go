package stores

import (
	"partyhaan/backend/models"
	"partyhaan/backend/stores/mariadb"
)

//UserStore ...
type UserStore struct {
}

func NewUserStore() *UserStore {
	return &UserStore{}
}
func (us *UserStore) Find(id int64) (models.User, error) {
	mdbUesr := mariadb.NewUser()
	return mdbUesr.Find(id)
}
func (us *UserStore) FindByParty(id int64, offset int64, limit int64) ([]models.User, error) {
	mdbUser := mariadb.NewUser()
	return mdbUser.FindByParty(id, offset, limit)
}
func (us *UserStore) Create(user models.User) (int64, error) {
	mdbUser := mariadb.NewUser()
	return mdbUser.Create(user)
}
func (us *UserStore) Update(id int64, user models.User) error {
	mdbUser := mariadb.NewUser()
	return mdbUser.Update(id, user)
}
func (us *UserStore) Delete(id int64) error {
	mdbUser := mariadb.NewUser()
	return mdbUser.Delete(id)
}
