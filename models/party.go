package models

//Party ...
type Party struct {
	Base
	Cover   string
	Name    string
	Limited int64
	Join    int64
	Owner   int64
}
