package models

type SenderID struct {
	Id          uint   `json:"id" gorm:"primaryKey"`
	Title       string `json:"title" gorm:"type:varchar(10);unique"`
	Description string `json:"description"`
	Base
}
