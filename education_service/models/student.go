package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Name   string  `json:"name"`
	Group  string  `json:"group"`
	Grades []Grade `json:"grades" gorm:"foreignKey:StudentID"`
}

type Grade struct {
	gorm.Model
	Subject   string  `json:"subject"`
	Value     float64 `json:"value"`
	Semester  int     `json:"semester"`
	StudentID uint
}
