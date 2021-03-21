package models

import "time"

//Base ...
type Base struct {
	ID      int64     `json:"id"`
	Created time.Time `json:"-"`
	Updated time.Time `json:"-"`
}
