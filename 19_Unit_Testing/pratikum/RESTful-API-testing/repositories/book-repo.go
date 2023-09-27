package repositories

import (
	"errors"
	"restful-api-testing/models"

	"gorm.io/gorm"
)
type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{
		db: db,
	}
}


func (br *BookRepository) SelectBooks() ([]models.Book, error) {
	var books []models.Book

	tx := br.db.Order("created_at desc").Find(&books)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return books, nil
}

func (br *BookRepository) SelectBookById(Id uint) (models.Book, error) {
	var book models.Book
	tx := br.db.First(&book, Id)
	if tx.Error != nil {
		return models.Book{}, tx.Error
	}

	return book, nil
}

func (br *BookRepository) InsertBook(book models.Book) error {
	tx := br.db.Create(&book)
	if tx.Error != nil {
		return errors.New("Insert book failed")
	}
	return nil
}

func (br *BookRepository) DeleteBook(Id uint) error {
	tx := br.db.Delete(&models.Book{}, Id)
	if tx.Error != nil {
		return errors.New("Delete book failed")
	}
	return nil
}

func (br *BookRepository) UpdateBook(book models.Book) error {
	updatedBook := models.Book{
		Title:     book.Title,
		Author:    book.Author,
		Publisher: book.Publisher,
	}

	tx := br.db.Model(&models.Book{}).Where("id = ?", book.ID).Updates(updatedBook)
	if tx.Error != nil {
		return errors.New("Update book failed")
	}

	if tx.RowsAffected == 0 {
		return errors.New("Book not found")
	}

	return nil
}
