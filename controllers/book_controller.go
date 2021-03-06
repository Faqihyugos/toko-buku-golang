package controllers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
	domain "toko/domain/model"
	"toko/services"
	"toko/utils"

	"github.com/gin-gonic/gin"
)

type bookController struct {
	BookService domain.IBookService
}

const (
	BOOK_LIST_PATH      = "/book/list"
	BOOK_CREATE_PATH    = "/book/add"
	BOOK_GET_BY_ID_PATH = "/book/:id"
	BOOK_UPDATE_PATH    = "/book/:id"
	BOOK_DELETE_PATH    = "/book/:id"
	ADD_STOCK_BOOK_PATH = "/book/:id/stock"
)

func NewBookController(db *sql.DB, r *gin.RouterGroup) {
	Controller := bookController{BookService: services.NewBookService(db)}
	r.GET(BOOK_LIST_PATH, Controller.lstBook)
	r.POST(BOOK_CREATE_PATH, Controller.AddBook)
	r.GET(BOOK_GET_BY_ID_PATH, Controller.GetBookById)
	r.PUT(ADD_STOCK_BOOK_PATH, Controller.addStockBook)
	r.PUT(BOOK_UPDATE_PATH, Controller.UpdateBook)

	r.DELETE(BOOK_DELETE_PATH, Controller.DeleteBook)
}

func (b *bookController) lstBook(c *gin.Context) {
	books, err := b.BookService.FindBook()
	fmt.Print("err", err)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, utils.Response(http.StatusOK, "All Get Data Success", books))
}

func (b *bookController) AddBook(c *gin.Context) {
	var book domain.Book
	err := c.ShouldBindJSON(&book)
	if err != nil {
		theErr := utils.NewUnprocessibleEntityError("invalid json body")
		c.JSON(theErr.Status(), theErr)
		return
	}

	newBook, errCreate := b.BookService.CreateBook(&book)
	if errCreate != nil {
		c.JSON(http.StatusInternalServerError, utils.NewInternalServerError(errCreate.Error()))
		return
	}
	c.JSON(http.StatusCreated, utils.Response(http.StatusCreated, "Book create successfully", newBook))
}

func (b *bookController) GetBookById(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		log.Println("Failed to convert to int")
		c.JSON(http.StatusInternalServerError, utils.NewInternalServerError("Internal server Error"))
	}
	book, er := b.BookService.FindBookById(id)
	if er != nil {
		log.Println(er.Error())
		c.JSON(http.StatusNotFound, utils.NewNotFoundError("book not found"))
	} else {
		c.JSON(http.StatusOK, utils.Response(http.StatusOK, "OK", book))
	}
}

func (b *bookController) addStockBook(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		log.Println("Failed to converted to int")
		c.JSON(http.StatusInternalServerError, utils.NewInternalServerError("Internal Server Error"))
	}
	var stock domain.Book
	errBind := c.ShouldBindJSON(&stock)
	if errBind != nil {
		theErr := utils.NewUnprocessibleEntityError("invalid json body")
		c.JSON(theErr.Status(), theErr)
		return
	}

	errStock := b.BookService.AddStock(stock.Stock, id)
	if errStock != nil {
		c.JSON(http.StatusInternalServerError, utils.NewInternalServerError("Internal Server Error"))
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 200, "message": "Add stock book successfully"})
	}
}

func (b *bookController) UpdateBook(c *gin.Context) {
	param := c.Param("id")
	id, errparse := strconv.Atoi(param)
	if errparse != nil {
		log.Println("Failed to converted to int")
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "Internal Server Error"})
	}

	var book domain.Book
	err := c.ShouldBindJSON(&book)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, utils.NewUnprocessibleEntityError("Invalid JSON Body"))
	}

	newBook, error := b.BookService.UpdateBook(&book, id)
	if err != nil {
		log.Println(error)
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "Internal Server Error"})
	}
	c.JSON(http.StatusOK, utils.Response(http.StatusOK, "Book updated successfully", newBook))
}

func (b *bookController) DeleteBook(c *gin.Context) {
	// Ambil id dari request
	param := c.Param("id")
	// parse yang tadi string ke int
	id, err := strconv.Atoi(param)
	if err != nil {
		log.Println("Failed to converted to int")
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "Internal Server Error"})
	}
	result, err := b.BookService.DeleteBook(id)
	log.Println("rows:", result)
	if err != nil {
		c.JSON(http.StatusNotFound, utils.NewNotFoundError("Book not found"))
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 200, "message": "Data deleted successfully", "data": result})
	}
}
