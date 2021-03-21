package models

//Join ...
type Join struct {
	Base
	PartyID int64 `json:"party_id"`
	UserID  int64 `json:"user_id"`
}
