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

// GET /books
// Find all books
// @Success 200 {string} FindBooks
// @Router /example/FindBooks [get]

func FindBooks(c *gin.Context) {
	var books []models.Book
	models.DB.Find(&books)

	c.JSON(http.StatusOK, gin.H{"data": books})
}

// GET /books/:id
// Find a book
// @Success 200 {string} FindBook
// @Router /example/FindBook [get]
func FindBook(c *gin.Context) {
	// Get model if exist
	var book models.Book
	models.DB.Where("id = ?", c.Param("id")).First(&book)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// POST /books
// Create new book
// @Success 200 {string} CreateBook
func CreateBook(c *gin.Context) {
	// Validate input
	var input CreateBookInput
	c.ShouldBindJSON(&input)
	// Create book
	book := models.Book{Title: input.Title, Author: input.Author}
	models.DB.Create(&book)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// @Router /example/CreateBook [post]

// PATCH /books/:id
// Update a book
// @Success 200 {string} UpdateBook
// @Router /example/UpdateBook [patch]
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

// DELETE /books/:id
// Delete a book
// @Success 200 {string} DeleteBook
// @Router /example/DeleteBook [delete]
func DeleteBook(c *gin.Context) {
	// Get model if exist
	var book models.Book
	models.DB.Where("id = ?", c.Param("id")).First(&book)

	models.DB.Delete(&book)

	c.JSON(http.StatusOK, gin.H{"data": true})

}
