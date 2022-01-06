package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Firstname string `json:"firstname"`
	Fastname  string `json:"lastname"`
	Age       uint32 `json:"age"`
	Id        string `json:"id"`
}
