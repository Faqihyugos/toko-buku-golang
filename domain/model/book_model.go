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
}

type IBookService interface {
	FindBook() ([]*Book, utils.MessageErr)
	CreateBook(book *Book) (*Book, error)
}
