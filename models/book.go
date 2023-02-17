package models

import (
	"github.com/jinzhu/gorm"
)

type Book struct {
	gorm.Model `json:"gorm_._model"`
	Title      string `json:"Title" json:"title,omitempty" json:"title,omitempty"`
	Author     string `json:"Author" json:"author,omitempty" json:"author,omitempty"`
}
