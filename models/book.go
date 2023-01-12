package models

type Book struct {
	ID     uint   `json:"" gorm:"primary_key"`
	Title  string `json:"Title"`
	Author string `json:"Author"`
}
