package controllers

import (
	"go-auth/database"
	"go-auth/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func List_department(c *gin.Context) {

	var Site_departments []models.Site_department

	if err := database.DB.Find(&Site_departments).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": "error", "data": "not found"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": "error", "data": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   Site_departments,
	})

}

func Show_department(c *gin.Context) {

	var Site_departments []models.Site_department
	var id = c.Param("id")

	if err := database.DB.First(&Site_departments, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": "error", "data": "not found"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": "error", "data": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   Site_departments,
	})

}

func Create_department(c *gin.Context) {

	var Site_departments models.Site_department

	if err := c.ShouldBindJSON(&Site_departments); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "error", "data": err.Error()})
		return
	}

	if err := database.DB.Create(&Site_departments).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": "error", "data": "not found"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": "error", "data": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Created New User",
	})

}

func Update_department(c *gin.Context) {

	var Site_departments models.Site_department
	id := c.Param("id")

	if err := c.ShouldBindJSON(&Site_departments); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "error", "data": err.Error()})
		return
	}

	if database.DB.Model(&Site_departments).Where("id = ?", id).Updates(&Site_departments).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Update User"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Update User",
	})

}
