package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model

	ID        int    `json:"ID" gorm: "primary_key";"auto_increment"`
	Name      string `json:"Name"`
	LastName  string `json:"LastName"`
	Address   string `json:"Address"`
	Telephone string `json:"Telephone"`
}
