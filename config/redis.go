package config

//Redis ...
type Redis struct {
	Addr     string `json:"addr"`
	Password string `json:"password"`
	DB       string `json:"db"`
}
