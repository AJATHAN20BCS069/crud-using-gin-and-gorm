package filters

import (
	"crud-demo/app/models"
	"crud-demo/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*// Filter users
func FilterUsersByName(c *gin.Context) {
    var users []models.User

    // Define a struct to hold the request body
    var requestBody struct {
        Name string `json:"name"`
    }

    // Bind JSON data from request body to the struct
    if err := c.BindJSON(&requestBody); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid JSON format in request body"})
        return
    }

    // Extract the name from the request body
    name := requestBody.Name

    // Check if name parameter is provided
    if name == "" {
        c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Name parameter is required"})
        return
    }

    db := database.GetDB()
    db.Where("name = ?", name).Find(&users)

    // Check if any users were found
    if len(users) == 0 {
        c.JSON(http.StatusNotFound, gin.H{"status": "error", "message": "No users found with the provided name"})
        return
    }

    c.JSON(http.StatusOK, users)
}*/

// Apply applies the given function to each element of the slice
func Apply(slice []models.User, function func(models.User) bool) []models.User {
    var result []models.User
    for _, elem := range slice {
        if function(elem) {
            result = append(result, elem)
        }
    }
    return result
}

func FilterUsersByName(c *gin.Context) {
    var users []models.User

    // Define a struct to hold the request body
    var requestBody struct {
        Name string `json:"name"`
    }

    // Bind JSON data from request body to the struct
    if err := c.BindJSON(&requestBody); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid JSON format in request body"})
        return
    }

    // Extract the name from the request body
    name := requestBody.Name

    // Check if name parameter is provided
    if name == "" {
        c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Name parameter is required"})
        return
    }

    db := database.GetDB()
    db.Find(&users)

    // Define a filtering function
    filterByName := func(user models.User) bool {
        return user.Name == name
    }

    // Apply the filtering function to the list of users
    filteredUsers := Apply(users, filterByName)

    // Check if any users were found
    if len(filteredUsers) == 0 {
        c.JSON(http.StatusNotFound, gin.H{"status": "error", "message": "No users found with the provided name"})
        return
    }

    c.JSON(http.StatusOK, filteredUsers)
}
