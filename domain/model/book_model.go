package model

import "toko/utils"

type Book struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"desc"`
	Year        int    `json:"year"`
	Pages       int    `json:"pages"`
	Language    string `json:"language"`
	Publisher   string `json:"publisher"`
	Price       int    `json:"price"`
	Stock       int    `json:"stock"`
}

type IBookRepository interface {
	Find() ([]*Book, utils.MessageErr)
	Create(book *Book) (*Book, error)
	FindById(id int) (*Book, error)
	Update(book *Book) (*Book, error)
	Delete(id int) (int64, error)
	UpdateStock(stock, id int) error
}

type IBookService interface {
	FindBook() ([]*Book, utils.MessageErr)
	CreateBook(book *Book) (*Book, error)
	FindBookById(id int) (*Book, error)
	UpdateBook(book *Book, id int) (*Book, error)
	DeleteBook(id int) (int64, error)
	AddStock(stock, id int) error
}
