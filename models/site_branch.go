package models

import "time"

type Site_branch struct {
	ID        uint      `gorm:"primaryKey; autoIncrement" json:"id"`
	Name      string    `json:"Name" binding:"required"`
	Address   string    `json:"Address" binding:"required"`
	Status    string    `json:"Status" binding:"required"`
	CreatedAt time.Time `gorm:"<-:create" json:"CreatedAt" binding:"required"`
	UpdatedAt time.Time `gorm:"<-:update" json:"UpdatedAt" binding:"required"`
	Users     []User
}
