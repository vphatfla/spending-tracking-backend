package model

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Password []byte `json:"password,omitempty"` // Use []byte for varbinary fields
}
