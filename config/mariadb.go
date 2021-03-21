package config

//MariaDB ...
type MariaDB struct {
	DSNQuery string `json:"dsn_query"`
	DSNExec  string `json:"dsn_exec"`
}
