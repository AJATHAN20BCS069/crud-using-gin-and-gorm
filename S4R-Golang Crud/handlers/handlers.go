package handlers

import (
	"crud-demo/app/models"
	"crud-demo/database"

	"github.com/gin-gonic/gin"
)

//create user
func CreateUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)

	db := database.GetDB()
	var existingUser models.User
	if err := db.First(&existingUser, "email = ?", user.Email).Error; err == nil {

		// User already exists, return an error response
		c.JSON(400, gin.H{"status": "error", "message": "User already exists"})
		return
	}

	// User does not exist, create a new user
	db.Create(&user)
	c.JSON(201, gin.H{"status": "success", "message": "User created successfully", "user": user})
}

//Get user
func GetUsers(c *gin.Context) {
	var users []models.User

	db := database.GetDB()
	db.Find(&users)

	c.JSON(200, users)
}

//Get user by id
func GetUserByID(c *gin.Context) {
	var user models.User
	id := c.Param("id")

	db := database.GetDB() // Declare db variable here

	if err := db.First(&user, id).Error; err != nil {
		c.JSON(404, gin.H{"status": "error", "message": "User not found"})
		return
	}

	c.JSON(200, user)
}

//update the user
func UpdateUser(c *gin.Context) {
	var user models.User
	id := c.Param("id")

	db := database.GetDB() // Declare db variable here

	if err := db.First(&user, id).Error; err != nil {
		c.JSON(404, gin.H{"status": "error", "message": "User not found"})
		return
	}

	c.BindJSON(&user)
	db.Save(&user)

	c.JSON(200, gin.H{"status": "success", "message": "User updated successfully", "user": user})
}

//delete the user
func DeleteUser(c *gin.Context) {
	var user models.User
	id := c.Param("id")

	db := database.GetDB() // Declare db variable here

	if err := db.First(&user, id).Error; err != nil {
		c.JSON(404, gin.H{"status": "error", "message": "User not found"})
		return
	}

	db.Delete(&user)
	c.JSON(200, gin.H{"status": "success", "message": "User deleted successfully"})
}
