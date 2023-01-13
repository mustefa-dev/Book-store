package controllers

import (
	models "Book-store/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UpdateBookInput struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

type CreateBookInput struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

func FindBooksByAuthorAndTitle(c *gin.Context) {
	var books []models.Book
	author := c.Query("author")
	title := c.Query("title")
	query := models.DB.Table("books")
	if title != "" {
		query = query.Where("title LIKE ?", title+"%")
	}
	if author != "" {
		query = query.Where("author = ?", author)
	}
	query.Find(&books)
	c.JSON(http.StatusOK, gin.H{"data": books})
}

func FindBooks(c *gin.Context) {
	var books []models.Book
	models.DB.Find(&books)

	c.JSON(http.StatusOK, gin.H{"data": books})
}

func FindBook(c *gin.Context) {
	// Get model if exist
	startsWith := c.Query("startsWith")
	var books []models.Book
	models.DB.Where("title LIKE ?", startsWith+"T%").Find(&books)
	c.JSON(http.StatusOK, gin.H{"data": books})
}

func CreateBook(c *gin.Context) {
	// Validate input
	var input CreateBookInput
	c.ShouldBindJSON(&input)
	// Create book
	book := models.Book{Title: input.Title, Author: input.Author}
	models.DB.Create(&book)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

func UpdateBook(c *gin.Context) {
	// Get model if exist
	var book models.Book
	models.DB.Where("id = ?", c.Param("id")).First(&book)

	// Validate input
	var input UpdateBookInput
	c.ShouldBindJSON(&input)

	models.DB.Model(&book).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

func DeleteBook(c *gin.Context) {
	// Get model if exist
	var book models.Book
	models.DB.Where("id = ?", c.Param("id")).First(&book)

	models.DB.Delete(&book)

	c.JSON(http.StatusOK, gin.H{"data": true})

}
