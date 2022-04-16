package app

import (
	"fredmnpinto/go-microservices/controllers"
	"github.com/gin-gonic/gin"
)

func StartApp() {
	router := gin.Default()

	router.GET("/books", controllers.GetBooks)
	router.GET("/books/:id", controllers.GetBookById)

	router.POST("/books", controllers.CreateBook)
	router.PATCH("/books/checkout/:id/:quantity", controllers.CheckoutBook)

	err := router.Run("localhost:8080")

	if err != nil {
		panic(err)
	}
}
