package pagination

import (
	"crud-demo/app/models"
	"crud-demo/database"
	"math"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Get users with pagination
func GetUsersWithPagination(c *gin.Context) {
	var users []models.User
	var page int
	var limit int

	// Parse query parameters for pagination
	if pageStr := c.Query("page"); pageStr != "" {
		parsedPage, err := strconv.Atoi(pageStr)
		if err != nil {
			c.JSON(400, gin.H{"status": "error", "message": "Invalid page number"})
			return
		}
		page = parsedPage
	} else {
		page = 1 // Default page number
	}

	if limitStr := c.Query("limit"); limitStr != "" {
		parsedLimit, err := strconv.Atoi(limitStr)
		if err != nil {
			c.JSON(400, gin.H{"status": "error", "message": "Invalid limit"})
			return
		}
		limit = parsedLimit
	} else {
		limit = 5 // Default limit
	}

	offset := (page - 1) * limit

	db := database.GetDB()
	db.Limit(limit).Offset(offset).Find(&users)

	c.JSON(200, users)
}

// search users by email
func SearchUsersByEmail(c *gin.Context) {
	var users []models.User
	var requestBody struct {
		Email string `json:"email"`
	}

	// Bind JSON data from request body to struct
	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(400, gin.H{"status": "error", "message": "Invalid JSON format in request body"})
		return
	}

	// Check if email parameter is provided
	if requestBody.Email == "" {
		c.JSON(400, gin.H{"status": "error", "message": "Email parameter is required"})
		return
	}

	db := database.GetDB()
	db.Where("email = ?", requestBody.Email).Find(&users)

	// Check if any users were found
	if len(users) == 0 {
		c.JSON(404, gin.H{"status": "error", "message": "No users found with the provided email"})
		return
	}

	c.JSON(200, users)
}

/*//postman params email another method

func SearchUsersByEmail(c *gin.Context) {
    var users []models.User
    email := c.Query("email")

    // Check if email parameter is provided
    if email == "" {
        c.JSON(400, gin.H{"status": "error", "message": "Email parameter is required"})
        return
    }

    db := database.GetDB()
    db.Where("email = ?", email).Find(&users)

    // Check if any users were found
    if len(users) == 0 {
        c.JSON(404, gin.H{"status": "error", "message": "No users found with the provided email"})
        return
    }

    c.JSON(200, users)
}*/

// Search user email with page value
/*func GetUserByEmail(c *gin.Context) {
    var user models.User

    // Search user by email
    email := c.Query("email")
    if email == "" {
        c.JSON(400, gin.H{"status": "error", "message": "Email parameter is required"})
        return
    }

    db := database.GetDB()
    db.Where("email = ?", email).First(&user)

    if user.ID == 0 {
        c.JSON(404, gin.H{"status": "error", "message": "User not found"})
        return
    }

    // Calculate page based on user's position in the database
    var count int64
    db.Model(&models.User{}).Where("id < ?", user.ID).Count(&count)
    page := int(math.Ceil(float64(count+1) / 5)) // Each page has a limit of 5 users

    // Prepare response
    response := gin.H{
        "user":  user,
        "page":  page,
    }

    c.JSON(200, response)
}*/

func GetUserByEmail(c *gin.Context) {
	var user models.User

	// Define a struct to hold the request body
	var requestBody struct {
		Email string `json:"email"`
	}

	// Bind JSON data from request body to the struct
	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(400, gin.H{"status": "error", "message": "Invalid JSON format in request body"})
		return
	}

	// Extract the email from the request body
	email := requestBody.Email

	// Check if email parameter is provided
	if email == "" {
		c.JSON(400, gin.H{"status": "error", "message": "Email parameter is required"})
		return
	}

	db := database.GetDB()
	db.Where("email = ?", email).First(&user)

	if user.ID == 0 {
		c.JSON(404, gin.H{"status": "error", "message": "User not found"})
		return
	}

	// Calculate page based on user's position in the database
	var count int64
	db.Model(&models.User{}).Where("id < ?", user.ID).Count(&count)
	page := int(math.Ceil(float64(count+1) / 5)) // Each page has a limit of 5 users

	// Prepare response
	response := gin.H{
		"user": user,
		"page": page,
	}

	c.JSON(200, response)
}
