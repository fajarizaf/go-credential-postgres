package models

import "time"

type Site_role struct {
	ID                uint      `gorm:"primaryKey; autoIncrement" json:"id"`
	Role_Name         string    `json:"Role_Name" binding:"required"`
	Role_Type         string    `json:"Role_Type" binding:"required"`
	Site_departmentID uint      `json:"department_id" binding:"required"`
	CreatedAt         time.Time `gorm:"<-:create" json:"CreatedAt" binding:"required"`
	UpdatedAt         time.Time `gorm:"<-:update" json:"UpdatedAt" binding:"required"`
	Users             []User
}
