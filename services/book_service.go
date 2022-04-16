package services

import (
	"fmt"
	"fredmnpinto/go-microservices/domain"
)

/* Mockup Database */
var booksDb = map[string]domain.Book{
	"1": {Id: "1", Title: "In search of Lost Time", Author: "Marcel Proust", Quantity: 2},
	"2": {Id: "2", Title: "The Little Prince", Author: "Exupery", Quantity: 5},
}

var (
	NotFoundError = fmt.Errorf("resource not found")

	NotEnoughUnitsError = fmt.Errorf("there aren't enough units for such checkout")
	RepeatedValuesError = fmt.Errorf("book already exists")
)

/* Request handlers */

// GetBooks /* Returns all books on the database */
func GetBooks() []domain.Book {
	var allBooks []domain.Book

	for _, val := range booksDb {
		allBooks = append(allBooks, val)
	}

	return allBooks
}

// CreateBook /* Inserts a new book onto the database */
func CreateBook(book domain.Book) error {
	if dbHasBook(book.Id) {
		return RepeatedValuesError
	}

	booksDb[book.Id] = book

	return nil
}

// GetBookById /* Returns the book with the specified Id */
func GetBookById(id string) (domain.Book, error) {

	if !dbHasBook(id) {
		return domain.Book{}, NotFoundError
	}

	return booksDb[id], nil
}

// CheckoutBook /* Updates the quantity of items of the given book */
func CheckoutBook(id string, quantity int) (int, error) {
	book, err := GetBookById(id)

	if err != nil {
		return -1, err
	}

	if book.Quantity-quantity < 0 {
		return book.Quantity, NotEnoughUnitsError
	}

	book.Quantity -= quantity
	booksDb[book.Id] = book

	return book.Quantity, nil
}

/* Private functions */

// dbHasBook /* Basic contains() function */
func dbHasBook(id string) bool {
	_, found := booksDb[id]

	return found
}
