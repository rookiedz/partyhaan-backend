package models

//Authen ...
type Authen struct {
	Base
	Email    string `json:"email"`
	Password string `json:"password"`
	Salt     string `json:"-"`
	Hash     string `json:"-"`
	UserID   int64  `json:"-"`
}
