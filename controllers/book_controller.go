package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	domain "toko/domain/model"
	"toko/services"
	"toko/utils"

	"github.com/gin-gonic/gin"
)

type bookController struct {
	BookService domain.IBookService
}

const (
	BOOK_LIST_PATH   = "/book/list"
	BOOK_CREATE_PATH = "/book"
)

func NewBookController(db *sql.DB, r *gin.RouterGroup) {
	Controller := bookController{BookService: services.NewBookService(db)}
	r.GET(BOOK_LIST_PATH, Controller.lstBook)
	r.POST(BOOK_CREATE_PATH, Controller.AddBook)

}

func (b *bookController) lstBook(c *gin.Context) {
	books, err := b.BookService.FindBook()
	fmt.Print("err", err)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": "Ok", "data": books})
}

func (b *bookController) AddBook(c *gin.Context) {
	var book domain.Book
	err := c.ShouldBindJSON(&book)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, utils.NewUnprocessibleEntityError("invalid json body"))
		return
	}

	newBook, errCreate := b.BookService.CreateBook(&book)
	if errCreate != nil {
		c.JSON(http.StatusInternalServerError, utils.NewInternalServerError("Internal Server Error"))
	}
	c.JSON(http.StatusCreated, utils.Response(http.StatusCreated, "Book create successfully", newBook))
}
