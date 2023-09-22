package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID            uint           `gorm:"primaryKey" json:"id"`
	Site_branchID int            `json:"branch_id" binding:"required"`
	Site_roleID   uint           `json:"role_id"`
	User_Login    string         `json:"User_Login" binding:"required"`
	User_Password string         `json:"User_Password" binding:"required"`
	Vcode         string         `json:"Vcode"`
	Rtoken        string         `json:"Rtoken"`
	ContactName   string         `json:"ContactName" binding:"required"`
	Address       string         `json:"Address" `
	City          string         `json:"City"`
	State         string         `json:"State"`
	Phone         string         `json:"Phone"`
	Isactive      bool           `json:"isactive" binding:"required"`
	Email         string         `gorm:"unique;not null" json:"Email" binding:"required"`
	CreatedAt     time.Time      `gorm:"<-:create" json:"CreatedAt"`
	UpdatedAt     time.Time      `gorm:"<-:update" json:"UpdatedAt"`
	CompanyID     []*Corporation `gorm:"many2many:user_corporations;" json:"CompanyID"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(u.User_Password), 14)
	if err != nil {
		return err
	}
	u.User_Password = string(bytes)
	return nil
}

func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(u.User_Password), 14)
	if err != nil {
		return err
	}
	u.User_Password = string(bytes)
	return nil
}

// CheckPassword checks user password
// CheckPassword takes a string as a parameter and compares it to the user's encrypted password
// It returns an error if there is an issue comparing the passwords
func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.User_Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}
