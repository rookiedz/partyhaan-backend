package models

import "time"

//Base ...
type Base struct {
	ID        int64     `json:"id"`
	Deletable time.Time `json:"-"`
	Created   time.Time `json:"-"`
	Updated   time.Time `json:"-"`
	Deleted   time.Time `json:"-"`
}
