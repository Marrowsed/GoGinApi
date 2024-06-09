package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Name     string `json: "name"`
	Document string `json: "document"`
	Age      int    `json: "age"`
}

var Students []Student
