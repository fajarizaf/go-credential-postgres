package models

import "time"

type Site_department struct {
	ID         uint        `gorm:"primaryKey; autoIncrement" json:"id"`
	Status     string      `json:"Status" binding:"required"`
	Department string      `json:"Department" binding:"required"`
	CreatedAt  time.Time   `gorm:"<-:create" json:"CreatedAt" binding:"required"`
	UpdatedAt  time.Time   `gorm:"<-:update" json:"UpdatedAt" binding:"required"`
	Roles      []Site_role `json:"roles"`
}
