package main

import (
	"go-auth/database"

	auth "go-auth/auth"
	controller "go-auth/controllers"

	"github.com/gin-gonic/gin"
)

func main() {

	database.ConnectDatabase()
	database.MigrateDatabase()

	r := gin.Default()

	// endpotin User
	r.GET("/api/user", controller.List_user)
	r.GET("/api/user/:id", controller.Show_user)
	r.POST("/api/user", controller.Create_user)
	r.PUT("/api/user/:id", controller.Update_user)

	// Endpoint Company
	r.GET("/api/company", controller.List_company)
	r.GET("/api/company/:id", controller.Show_company)
	r.POST("/api/company", controller.Create_company)
	r.PUT("/api/company/:id", controller.Update_company)

	// Endpoint Site Role
	r.GET("/api/role", controller.List_role)
	r.GET("/api/role/:id", controller.Show_role)
	r.POST("/api/role", controller.Create_role)
	r.PUT("/api/role/:id", controller.Update_role)

	// Endpoint Site Department
	r.GET("/api/department", controller.List_department)
	r.GET("/api/department/:id", controller.Show_department)
	r.POST("/api/department", controller.Create_department)
	r.PUT("/api/department/:id", controller.Update_department)

	// Endpoint Authentication
	r.POST("/api/auth/login", auth.Login)
	r.POST("/api/auth/register", auth.Register)

	r.Run()

}
