package controllers

import (
	"fmt"
	"fredmnpinto/go-microservices/domain"
	"fredmnpinto/go-microservices/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetBooks(context *gin.Context) {
	allBooks := services.GetBooks()

	context.IndentedJSON(http.StatusOK, allBooks)
}

func CreateBook(context *gin.Context) {
	var newBook domain.Book

	if err := context.BindJSON(&newBook); err != nil {
		context.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	err := services.CreateBook(newBook)
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	context.IndentedJSON(http.StatusCreated, newBook)
}

func GetBookById(context *gin.Context) {
	id := context.Param("id")
	book, err := services.GetBookById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, err)
		return
	}

	context.IndentedJSON(http.StatusOK, book)
}

func CheckoutBook(context *gin.Context) {
	id := context.Param("id")

	quantity, qtyErr := strconv.Atoi(context.Param("quantity"))

	if qtyErr != nil {
		errErr := context.AbortWithError(http.StatusBadRequest, qtyErr)

		if errErr != nil {
			fmt.Print("Fuck: ", errErr)
		}

		return
	}

	newQuantity, serviceErr := services.CheckoutBook(id, quantity)

	if serviceErr != nil {
		cErr := context.AbortWithError(http.StatusBadRequest, serviceErr)

		if cErr != nil {
			panic(cErr)
		}

		return
	}

	context.IndentedJSON(http.StatusOK, gin.H{"newQuantity": newQuantity})
}
