package services

import (
	"database/sql"
	"fmt"
	domain "toko/domain/model"
	"toko/domain/repositories"
	"toko/utils"
)

type bookService struct {
	db       *sql.DB
	BookRepo domain.IBookRepository
}

func NewBookService(db *sql.DB) domain.IBookService {
	return &bookService{db: db, BookRepo: repositories.NewBookRepo(db)}
}

func (b *bookService) FindBook() ([]*domain.Book, utils.MessageErr) {
	books, err := b.BookRepo.Find()
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (b *bookService) CreateBook(book *domain.Book) (*domain.Book, error) {
	book, err := b.BookRepo.Create(book)
	fmt.Println("Service :", book)
	if err != nil {
		return nil, err
	}
	return book, nil
}
