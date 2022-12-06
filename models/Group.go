package models

type Group struct {
	Id          uint   `json:"id" gorm:"primaryKey"`
	Title       string `json:"title" gorm:"unique"`
	Description string `json:"description"`
	Base
}
