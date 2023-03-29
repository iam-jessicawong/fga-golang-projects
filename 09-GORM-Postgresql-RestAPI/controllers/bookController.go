package controllers

import (
	"fmt"
	"gorm_postgresql_restapi/database"
	"gorm_postgresql_restapi/models"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetBooks(ctx *gin.Context) {
	db := database.GetDB() // get db function from database package

	var books = []models.Book{} //variable to store data
	// get book from database
	if err := db.Find(&books).Error; err != nil {
		log.Fatal("Error getting books data:", err.Error())
	}

	ctx.JSON(http.StatusOK, books)
}

func GetBookById(ctx *gin.Context) {
	db := database.GetDB()
	bookID := ctx.Param("id")
	book := models.Book{}

	if err := db.First(&book, "id = ?", bookID).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("Book data with id %v is not found", bookID),
			"status":  "not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, book)
}

func CreateBook(ctx *gin.Context) {
	db := database.GetDB()
	var newBook models.Book

	// validate body request
	if err := ctx.ShouldBindJSON(&newBook); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "invalid inputs, please rechek your input",
		})
		return
	}

	// trim space from input
	newBook.Title = strings.TrimSpace(newBook.Title)
	newBook.Author = strings.TrimSpace(newBook.Author)
	newBook.Description = strings.TrimSpace(newBook.Description)

	if err := db.Create(&newBook).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "successfully add new book",
		"data":    newBook,
	})

}

func UpdateBook(ctx *gin.Context) {
	db := database.GetDB()
	bookID := ctx.Param("id")
	book := models.Book{}

	// check if data is exist
	if err := db.Where("id = ?", bookID).First(&book).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("the book with id %v that you're trying to update is not exist", bookID),
			"status":  "not found",
		})
		return
	}

	// validate input
	updatedBook := models.Book{}
	err := ctx.ShouldBindJSON(&updatedBook)
	// trim space from input
	updatedBook.Title = strings.TrimSpace(updatedBook.Title)
	updatedBook.Author = strings.TrimSpace(updatedBook.Author)
	updatedBook.Description = strings.TrimSpace(updatedBook.Description)

	if err != nil || len(updatedBook.Title) < 2 || len(updatedBook.Author) < 2 || len(updatedBook.Description) < 5 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "invalid inputs, please recheck your inputs",
		})
		return
	}

	result := db.Model(&book).Updates(updatedBook)
	if result.Error != nil || result.RowsAffected == 0 {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("book with id %v successfully updated", bookID),
	})
}

func DeleteBook(ctx *gin.Context) {
	db := database.GetDB()
	bookID := ctx.Param("id")
	book := models.Book{}

	// check if book is exist
	if err := db.Where("id = ?", bookID).First(&book).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("the book you're trying to delete with id %v is not exist", bookID),
			"status":  "not found",
		})
		return
	}

	if err := db.Delete(&book).Error; err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("book with id %v has been successfully deleted", bookID),
	})
}
