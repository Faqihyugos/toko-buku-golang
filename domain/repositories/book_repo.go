package repositories

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	domain "toko/domain/model"
	"toko/utils"
)

type bookRepo struct {
	db *sql.DB
}

func NewBookRepo(db *sql.DB) domain.IBookRepository {
	return &bookRepo{db: db}
}

func (c bookRepo) Find() ([]*domain.Book, utils.MessageErr) {
	// Membuat object slice category
	books := make([]*domain.Book, 0)
	//defer a.db.Close()

	// Untuk format query
	query := fmt.Sprintf(`SELECT id, title, description, year, pages, language, publisher, price, stock FROM book`)

	// Eksekusi query
	rows, err := c.db.Query(query)
	if err != nil {
		return nil, utils.ParserError(err)
	}
	defer rows.Close()

	for rows.Next() {
		book := &domain.Book{}
		getError := rows.Scan(&book.Id, &book.Title, &book.Description, &book.Year, &book.Pages, &book.Language,
			&book.Publisher, &book.Price, &book.Stock)
		if err != nil {
			return nil, utils.NewInternalServerError(fmt.Sprintf("Error when trying to get message: %s", getError.Error()))
		}
		books = append(books, book)
	}
	if len(books) == 0 {
		return nil, utils.NewNotFoundError("no records found")
	}
	return books, nil
}

func (c *bookRepo) Create(book *domain.Book) (*domain.Book, error) {
	query := fmt.Sprintf(`INSERT INTO book(title, description, year, pages, language, publisher, price, stock) VALUES(?,?,?,?,?,?,?,?)`)
	result, err := c.db.Exec(
		query,
		&book.Title,
		&book.Description,
		&book.Year,
		&book.Pages,
		&book.Language,
		&book.Publisher,
		&book.Price,
		&book.Stock,
	)

	if err != nil {
		s := strings.Split(err.Error(), ":")
		log.Println(s[1])
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, utils.NewInternalServerError(fmt.Sprintf("Error when trying to save massage: %s", err.Error()))
	}

	book.Id = int(id)

	return book, nil
}
