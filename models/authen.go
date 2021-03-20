package models

//Authen ...
type Authen struct {
	Base
	Email    string
	Password string
	Salt     string
	Hash     string
}
