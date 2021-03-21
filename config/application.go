package config

//Application ...
type Application struct {
	Port    int64  `json:"port"`
	Hash256 string `json:"hash256"`
	Hash512 string `json:"hash512"`
}
