package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Book struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"desc"`
}

var bookDatas = []Book{
	{
		ID:          1,
		Title:       "Harry Potter and the Sorcerer's Stone",
		Author:      "J.K. Rowling",
		Description: "Its about Harry!",
	}, {
		ID:          2,
		Title:       "And Then There Were None",
		Author:      "Agatha Christie",
		Description: "Its a Mystery!",
	},
	{
		ID:          3,
		Title:       "The Hobbits",
		Author:      "J. R. R. Tolkieng",
		Description: "Hooooooobiiiiiiiiiiiittts!",
	},
}

func GetBooks(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, bookDatas)
	return
}

func GetBook(ctx *gin.Context) {
	var bookData Book
	status := false

	// Check params
	bookIdString := ctx.Param("id")
	bookId, err := strconv.Atoi(bookIdString)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	// Find data
	for _, book := range bookDatas {
		if book.ID == bookId {
			status = true
			bookData = book
			break
		}
	}

	if !status {
		response := fmt.Sprintf("Book with ID %d not found", bookId)
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": response,
		})
		return
	}

	// Success get single data
	ctx.JSON(http.StatusOK, bookData)
	return
}

func AddBook(ctx *gin.Context) {
	var newBook Book
	if err := ctx.ShouldBindJSON(&newBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// New id is from the latest book id + 1
	getLatestBook := bookDatas[len(bookDatas)-1]
	newBook.ID = getLatestBook.ID + 1
	bookDatas = append(bookDatas, newBook)

	// response := fmt.Sprintf("Success add book %s", newBook.Title)
	ctx.JSON(http.StatusCreated, "Created")
	return
}

func UpdateBook(ctx *gin.Context) {
	var updatedBook Book
	status := false

	// Check Params (should be integer)
	bookIdString := ctx.Param("id")
	bookId, err := strconv.Atoi(bookIdString)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	// Check bind json
	if err := ctx.ShouldBindJSON(&updatedBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// Find and Change Datas
	for idx, book := range bookDatas {
		if book.ID == bookId {
			status = true
			updatedBook.ID = book.ID
			bookDatas[idx] = updatedBook
		}
	}

	if !status {
		response := fmt.Sprintf("Book with ID %d not found", bookId)
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": response,
		})
		return
	}

	// response := fmt.Sprintf("Update data for ID %d success.", bookId)
	ctx.JSON(http.StatusOK, "Updated")
	return
}
func DeleteBook(ctx *gin.Context) {
	var bookIndex int
	status := false

	// Check Params (should be integer)
	bookIdString := ctx.Param("id")
	bookId, err := strconv.Atoi(bookIdString)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	// Find and get index of book
	for idx, book := range bookDatas {
		if book.ID == bookId {
			status = true
			bookIndex = idx
			break
		}
	}
	if !status {
		response := fmt.Sprintf("Book with ID %d not found", bookId)
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": response,
		})
		return
	}

	bookDatas = append(bookDatas[:bookIndex], bookDatas[bookIndex+1:]...)

	// response := fmt.Sprintf("Delete data for ID %d success.", bookId)
	ctx.JSON(http.StatusOK, "Deleted")
	return

}
