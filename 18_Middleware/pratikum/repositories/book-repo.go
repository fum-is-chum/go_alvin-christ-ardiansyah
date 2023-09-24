package repositories

import (
	"errors"
	"latihan/config"
	"latihan/models"
)

func SelectBooks() ([]models.Book, error) {
	var books []models.Book

	tx := config.DB.Order("created_at desc").Find(&books)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return books, nil
}

func SelectBookById(Id uint) (models.Book, error) {
	var book models.Book
	tx := config.DB.First(&book, Id)
	if tx.Error != nil {
		return models.Book{}, tx.Error
	}

	return book, nil
}

func InsertBook(book models.Book) error {
	tx := config.DB.Create(&book)
	if tx.Error != nil {
		return errors.New("Insert book failed")
	}
	return nil
}

func DeleteBook(Id uint) error {
	tx := config.DB.Delete(&models.Book{}, Id)
	if tx.Error != nil {
		return errors.New("Delete book failed")
	}
	return nil
}

func UpdateBook(book models.Book) error {
	updatedBook := models.Book{
		Title:     book.Title,
		Author:    book.Author,
		Publisher: book.Publisher,
	}

	tx := config.DB.Model(&models.Book{}).Where("id = ?", book.ID).Updates(updatedBook)
	if tx.Error != nil {
		return errors.New("Update book failed")
	}

	if tx.RowsAffected == 0 {
		return errors.New("Book not found")
	}

	return nil
}
