package stores

import (
	"partyhaan/backend/models"
	"partyhaan/backend/stores/mariadb"
)

//JoinStore ...
type JoinStore struct{}

func NewJoinStore() *JoinStore {
	return &JoinStore{}
}

func (js *JoinStore) Create(join models.Join) (int64, error) {
	mdbJoin := mariadb.NewJoin()
	return mdbJoin.Create(join)
}
func (js *JoinStore) Delete(id int64) error {
	mdbJoin := mariadb.NewJoin()
	return mdbJoin.Delete(id)
}
