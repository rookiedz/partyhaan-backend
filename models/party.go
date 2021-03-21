package models

//Party ...
type Party struct {
	Base
	Cover string `json:"cover,omitempty"`
	Name  string `json:"name,omitempty"`
	Limit int64  `json:"limit,omitempty"`
	Join  int64  `json:"join,omitempty"`
	Owner int64  `json:"owner,omitempty"`
}
