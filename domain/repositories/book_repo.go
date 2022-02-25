package repositories

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"strings"
	domain "toko/domain/model"
	"toko/utils"
)

type bookRepo struct {
	db *sql.DB
}

// NewBookRepo is a constructor
func NewBookRepo(db *sql.DB) domain.IBookRepository {
	return &bookRepo{db: db}
}

func (c bookRepo) Find() ([]*domain.Book, utils.MessageErr) {
	// Membuat object slice category
	books := make([]*domain.Book, 0)
	// Untuk format query
	query := fmt.Sprintf(`SELECT id, title, description, year, pages, language, publisher, price, stock, purchase_amount FROM book`)

	// Eksekusi query
	rows, err := c.db.Query(query)
	if err != nil {
		return nil, utils.ParserError(err)
	}
	defer rows.Close()

	for rows.Next() {
		book := &domain.Book{}
		getError := rows.Scan(&book.Id, &book.Title, &book.Description, &book.Year, &book.Pages, &book.Language,
			&book.Publisher, &book.Price, &book.Stock, &book.PurchaseAmount)
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
	result, err := c.db.Exec(query, &book.Title, &book.Description, &book.Year, &book.Pages, &book.Language,
		&book.Publisher, &book.Price, &book.Stock)
	if err != nil {
		s := strings.Split(err.Error(), ":")
		log.Println(s[1])
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, utils.ParserError(err)
	}

	book.Id = int(id)

	return book, nil
}

func (c bookRepo) FindById(id int) (*domain.Book, error) {
	book := new(domain.Book)

	query := fmt.Sprintf(`SELECT id, title, description, year, pages,language, publisher, price, stock FROM book WHERE id=?`)
	if getError := c.db.QueryRow(query, id).
		Scan(&book.Id, &book.Title, &book.Description, &book.Year, &book.Pages,
			&book.Language, &book.Publisher, &book.Price, &book.Stock); getError != nil {
		fmt.Println("this is the error man: ", getError)
		return nil, getError
	}
	return book, nil
}

func (c bookRepo) Update(book *domain.Book) (*domain.Book, error) {
	query := fmt.Sprintf("UPDATE book SET title = ?, description = ?, year = ?, pages = ?, language = ?, publisher = ?, price = ?, stock = ? WHERE id = ?")
	_, updateErr := c.db.Exec(query, &book.Title, &book.Description, &book.Year, &book.Pages,
		&book.Language, &book.Publisher, &book.Price, &book.Stock, &book.Id)
	if updateErr != nil {
		s := strings.Split(updateErr.Error(), ":")
		log.Println(s[1])
		if updateErr != nil {
			return nil, updateErr
		}
	}

	return book, nil
}

func (c bookRepo) Delete(id int) (int64, utils.MessageErr) {
	query := fmt.Sprintf("DELETE FROM book WHERE id = ?")
	result, err := c.db.Exec(query, id)
	if err != nil {
		return 0, utils.ParserError(err)
	}
	RowsAffected, errRows := result.RowsAffected()
	if errRows != nil {
		return 0, utils.ParserError(errRows)
	}
	return RowsAffected, nil
}

func (c bookRepo) UpdateStock(book *domain.Book) (*domain.Book, error) {
	fmt.Println("In Book Repo : ", &book.Stock, &book.Id)
	query := fmt.Sprintf("UPDATE book SET stock = ? WHERE id = ?")
	_, updateErr := c.db.Exec(query, &book.Stock, strconv.Itoa(book.Id))
	if updateErr != nil {
		return nil, utils.ParserError(updateErr)
	}
	return book, nil
}

func (c *bookRepo) UpdatePurchaseAmount(book *domain.Book) (*domain.Book, utils.MessageErr) {
	fmt.Println("In Book Repo : ", &book.PurchaseAmount, &book.Id)
	query := fmt.Sprintf("UPDATE book SET purchase_amount = ? WHERE id = ?")
	// Eksekusi query
	_, updateErr := c.db.Exec(query, &book.PurchaseAmount, strconv.Itoa(book.Id))
	if updateErr != nil {
		return nil, utils.ParserError(updateErr)
	}
	return book, nil
}
