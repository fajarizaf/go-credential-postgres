package controllers

import (
	"go-auth/database"
	"go-auth/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func List_company(c *gin.Context) {

	var companys []models.Corporation

	if err := database.DB.Find(&companys).Error; err != nil {
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
		"data":   companys,
	})

}

func Show_company(c *gin.Context) {

	var companys []models.Corporation
	var id = c.Param("id")

	if err := database.DB.First(&companys, id).Error; err != nil {
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
		"data":   companys,
	})

}

func Create_company(c *gin.Context) {

	var companys models.Corporation

	if err := c.ShouldBindJSON(&companys); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "error", "data": err.Error()})
		return
	}

	if err := database.DB.Create(&companys).Error; err != nil {
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

func Update_company(c *gin.Context) {

	var companys models.Corporation
	id := c.Param("id")

	if err := c.ShouldBindJSON(&companys); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "error", "data": err.Error()})
		return
	}

	if database.DB.Model(&companys).Where("id = ?", id).Updates(&companys).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Update User"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Update User",
	})

}
