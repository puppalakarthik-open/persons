package models

import "gorm.io/gorm"

type person struct {
	Name      *string `jso:"name"`
	ID        uint    `gorm:"primary key; auto Increment;json:"id"`
	Gender    *string `json:"gender"`
	BloodType *string `json:"bloodtype"`
	Num       int32   `json:"num"`
}

func MigratePersons(db *gorm.DB) error {
	err := db.AutoMigrate(&person{})
	return err

}
