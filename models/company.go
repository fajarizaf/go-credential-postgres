package models

import (
	"time"

	"gorm.io/gorm"
)

type Corporation struct {
	ID          uint      `gorm:"primaryKey; autoIncrement" json:"id"`
	Parent_ID   uint      `json:"Parent_ID" binding:"required"`
	CompanyName string    `gorm:"unique;not null" json:"CompanyName" binding:"required"`
	Address     string    `json:"Address" binding:"required"`
	City        string    `json:"City" binding:"required"`
	State       string    `json:"State" binding:"required"`
	Phone       string    `json:"Phone" binding:"required"`
	Email       string    `gorm:"unique;not null" json:"Email" binding:"required"`
	CreatedAt   time.Time `json:"<-:create" binding:"required"`
	UpdatedAt   time.Time `json:"<-:update" binding:"required"`
	UserID      []*User   `gorm:"many2many:user_corporations;" json:"UserID"`
}

func (c *Corporation) BeforeCreate(tx *gorm.DB) (err error) {
	c.Parent_ID = 2
	return nil
}

func (c *Corporation) BeforeSave(tx *gorm.DB) (err error) {
	c.Parent_ID = 2
	return nil
}
