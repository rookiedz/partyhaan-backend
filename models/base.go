package models

import "time"

//Base ...
type Base struct {
	ID        int64
	Deletable time.Time
	Created   time.Time
	Updated   time.Time
	Deleted   time.Time
}
