package models

type User struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Email    string `json:"email"`
	password string `json:"password,omitempty"`
	Address  string `json:"address"`
}
