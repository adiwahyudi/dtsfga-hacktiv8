package routers

import (
	"challenge-2-simple-rest-api/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.GET("/books", controllers.GetBooks)
	router.GET("/book/:id", controllers.GetBook)
	router.POST("/book", controllers.AddBook)
	router.PUT("/book/:id", controllers.UpdateBook)
	router.DELETE("/book/:id", controllers.DeleteBook)

	return router
}
