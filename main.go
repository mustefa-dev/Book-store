package main

import (
	"Book-store/controllers"
	"Book-store/models"
	. "Book-store/models"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"       // swagger embed files
	"github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func main() {
	r := gin.Default()
	// Connect to database
	models.ConnectDatabase()

	// Routes
	r.GET("/findBooks", ErrorHandler(), controllers.FindBooks)
	r.GET("/findBook/:id", ErrorHandler(), controllers.FindBook)
	r.GET("/books/filter", ErrorHandler(), controllers.FindBooksByAuthorAndTitle)
	r.POST("/createBook", ErrorHandler(), controllers.CreateBook)
	r.POST("/UpdateBook/:id", ErrorHandler(), controllers.UpdateBook)
	r.DELETE("/DeleteBook/:id", ErrorHandler(), controllers.DeleteBook)
	r.Static("/swaggerui/", "cmd/api/swaggerui")

	// Serve the Swagger documentation
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	// Run the server
	err := r.Run()

	if err != nil {
		return
	}
}
