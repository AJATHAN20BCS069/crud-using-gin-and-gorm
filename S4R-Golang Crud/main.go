package main

import (
	"crud-demo/database"
	"crud-demo/filters"
	"crud-demo/handlers"
	"crud-demo/pagination"

	"github.com/gin-gonic/gin"
)

func main() {

	// To Initialize the database
	database.InitDB()

	// Create a new Gin router
	r := gin.Default()

	// Define routes
	r.POST("/users", handlers.CreateUser)
	r.GET("/users", handlers.GetUsers)
	r.GET("/users/:id", handlers.GetUserByID)
	r.PUT("/users/:id", handlers.UpdateUser)
	r.DELETE("/users/:id", handlers.DeleteUser)
	r.GET("/users/paginated", pagination.GetUsersWithPagination)
	r.POST("/get-users", pagination.GetUserByEmail)
	r.POST("/search-users", pagination.SearchUsersByEmail)
	r.POST("/filter-name", filters.FilterUsersByName)

	// Start the HTTP server
	r.Run(":3000")

}
