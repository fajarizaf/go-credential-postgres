package controllers

import (
	"go-auth/database"
	"go-auth/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func List_role(c *gin.Context) {

	var Site_roles []models.Site_role

	if err := database.DB.Find(&Site_roles).Error; err != nil {
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
		"data":   Site_roles,
	})

}

func Show_role(c *gin.Context) {

	var Site_roles []models.Site_role
	var id = c.Param("id")

	if err := database.DB.First(&Site_roles, id).Error; err != nil {
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
		"data":   Site_roles,
	})

}

func Create_role(c *gin.Context) {

	var Site_roles models.Site_role

	if err := c.ShouldBindJSON(&Site_roles); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "error", "data": err.Error()})
		return
	}

	if err := database.DB.Create(&Site_roles).Error; err != nil {
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

func Update_role(c *gin.Context) {

	var Site_roles models.Site_role
	id := c.Param("id")

	if err := c.ShouldBindJSON(&Site_roles); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "error", "data": err.Error()})
		return
	}

	if database.DB.Model(&Site_roles).Where("id = ?", id).Updates(&Site_roles).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Update User"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Update User",
	})

}
