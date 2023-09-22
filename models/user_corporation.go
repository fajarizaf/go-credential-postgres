package models

import "gorm.io/gorm"

type User_corporation struct {
	Corporation_id int
	User_id        int
}

func (corporation *User_corporation) Asign(DB *gorm.DB) error {
	result := DB.Create(&corporation)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
