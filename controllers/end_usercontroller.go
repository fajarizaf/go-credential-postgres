package controllers

import (
	"go-auth/database"
	"go-auth/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func List_user(c *gin.Context) {

	var users []models.User

	if err := database.DB.Find(&users).Error; err != nil {
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
		"data":   users,
	})

}

func Show_user(c *gin.Context) {

	var users []models.User
	var id = c.Param("id")

	if err := database.DB.First(&users, id).Error; err != nil {
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
		"data":   users,
	})

}

func Create_user(c *gin.Context) {

	var users models.User

	if err := c.ShouldBindJSON(&users); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "error", "data": err.Error()})
		return
	}

	if err := database.DB.Create(&users).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": "error", "data": "not found"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": "error", "data": err.Error()})
			return
		}
	}

	// mapping user to default company
	database.DB.Last(&users)
	usercorp := models.User_corporation{
		Corporation_id: 1,
		User_id:        int(users.ID),
	}

	usercorp.Asign(database.DB)

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Created New User",
	})

}

func Update_user(c *gin.Context) {

	var users models.User
	id := c.Param("id")

	if err := c.ShouldBindJSON(&users); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "error", "data": err.Error()})
		return
	}

	if database.DB.Model(&users).Where("id = ?", id).Updates(&users).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Update User"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Update User",
	})

}
