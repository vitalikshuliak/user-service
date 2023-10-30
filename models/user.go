package models

type User struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}
