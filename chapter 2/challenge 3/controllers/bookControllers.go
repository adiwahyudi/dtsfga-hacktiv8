package controllers

import (
	"chapter2-challenge3/database"
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

func GetBooks(ctx *gin.Context) {
	res := []Book{}
	statement := `SELECT * from book`

	rows, err := database.DB.Query(statement)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var books = Book{}
		err = rows.Scan(&books.ID, &books.Title, &books.Author, &books.Description)

		if err != nil {
			panic(err)
		}

		res = append(res, books)
	}

	ctx.JSON(http.StatusOK, res)
	return
}

func GetBook(ctx *gin.Context) {
	res := Book{}
	statement := `SELECT * from book WHERE id = $1`

	// Check params
	bookIdString := ctx.Param("id")
	bookId, err := strconv.Atoi(bookIdString)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	row := database.DB.QueryRow(statement, bookId)
	err = row.Scan(&res.ID, &res.Title, &res.Author, &res.Description)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Success get single data
	ctx.JSON(http.StatusOK, res)
	return
}

func AddBook(ctx *gin.Context) {
	newBook := Book{}
	statement := `INSERT INTO book (title, author, description) VALUES($1, $2, $3) Returning *`

	if err := ctx.ShouldBindJSON(&newBook); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	row := database.DB.QueryRow(statement, newBook.Title, newBook.Author, newBook.Description)
	if err := row.Scan(&newBook.ID, &newBook.Title, &newBook.Author, &newBook.Description); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, newBook)
	return
}

func UpdateBook(ctx *gin.Context) {
	updatedBook := Book{}
	statement := `UPDATE book set title=$2, author=$3, description=$4 where id=$1 RETURNING *`

	// Check Params (should be integer)
	bookIdString := ctx.Param("id")
	bookId, err := strconv.Atoi(bookIdString)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Check bind json
	if err := ctx.ShouldBindJSON(&updatedBook); err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	row := database.DB.QueryRow(statement, bookId, updatedBook.Title, updatedBook.Author, updatedBook.Description)
	err = row.Scan(&updatedBook.ID, &updatedBook.Title, &updatedBook.Author, &updatedBook.Description)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, updatedBook)
	return
}
func DeleteBook(ctx *gin.Context) {
	deletedBook := Book{}
	statement := `DELETE from book WHERE id=$1 RETURNING *`

	// Check Params (should be integer)
	bookIdString := ctx.Param("id")
	bookId, err := strconv.Atoi(bookIdString)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	row := database.DB.QueryRow(statement, bookId)
	if err := row.Scan(&deletedBook.ID, &deletedBook.Title, &deletedBook.Author, &deletedBook.Description); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, deletedBook)
	return

}
