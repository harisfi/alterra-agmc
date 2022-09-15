package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/harisfi/alterra-agmc/day3/models"
)

func GetAllBooks(c echo.Context) error {
	books := []models.Book{
		{
			IDModel:        models.IDModel{ID: 1},
			Title:          "Book A",
			Author:         "Book A",
			Publisher:      "Book A",
			TimestampModel: models.TimestampModel{},
		},
		{
			IDModel:        models.IDModel{ID: 2},
			Title:          "Book B",
			Author:         "Book B",
			Publisher:      "Book B",
			TimestampModel: models.TimestampModel{},
		},
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success get all books",
		"books":   books,
	})
}

func GetBookById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	book := models.Book{
		IDModel:        models.IDModel{ID: uint(id)},
		Title:          "Book A",
		Author:         "Book A",
		Publisher:      "Book A",
		TimestampModel: models.TimestampModel{},
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success get book by id",
		"book":    book,
	})
}

func CreateBook(c echo.Context) error {
	book := models.Book{}
	c.Bind(&book)
	book.ID = 3

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success create new book",
		"book":    book,
	})
}

func UpdateBookById(c echo.Context) error {
	book := models.Book{}
	c.Bind(&book)
	id, _ := strconv.Atoi(c.Param("id"))
	book.ID = uint(id)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success update book by id",
		"book":    book,
	})
}

func DeleteBookById(c echo.Context) error {
	book := models.Book{}
	id, _ := strconv.Atoi(c.Param("id"))
	book.ID = uint(id)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success delete book by id",
	})
}
