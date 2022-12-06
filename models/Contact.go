package models

type Contact struct {
	Id        uint   `json:"id" gorm:"primaryKey"`
	FirstName string `json:"firstName" gorm:"type:varchar(255)"`
	LastName  string `json:"lastName" gorm:"type:varchar(255)"`
	Email     string `json:"email" gorm:"type:varchar(500);unique"`
	Phone     string `json:"phone" gorm:"type:varchar(20);unique"`
	GroupID   uint   `json:"GroupID"`
	Group     Group  `json:"-" gorm:"foreignKey:GroupID"`
    Base
}
