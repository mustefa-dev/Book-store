package main

import (
	"Book-store/controllers"
	"Book-store/models"
	. "Book-store/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Connect to database
	models.ConnectDatabase()

	// Routes
	r.GET("/findBooks", ErrorHandler(), controllers.FindBooks)
	r.GET("/findBook/:id", ErrorHandler(), controllers.FindBook)
	r.POST("/createBook", ErrorHandler(), controllers.CreateBook)
	r.POST("/UpdateBook/:id/", ErrorHandler(), controllers.UpdateBook)
	r.DELETE("/DeleteBook/:id", ErrorHandler(), controllers.DeleteBook)

	// Run the server
	err := r.Run()
	if err != nil {
		return
	}
}
