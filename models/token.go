package models

import "time"

//Token ...
type Token struct {
	AccessToken string
	ExpiresIn   time.Time
}
