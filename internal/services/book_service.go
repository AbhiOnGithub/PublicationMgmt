package services

import (
	"github.com/abhiongithub/publicationmgmt/internal/models"
	"github.com/jinzhu/gorm"
)

type Service struct {
	DB *gorm.DB
}

func NewService(db *gorm.DB) *Service {
	return &Service{
		DB: db,
	}
}

type BookService interface {
	GetBookById(id uint) (models.Book, error)
	GetBookByTitle(title string) (models.Book, error)
	PostBook(book models.Book) (models.Book, error)
	UpdateBook(id uint, book models.Book) (models.Book, error)
	DeleteBook(id uint) error
	GetAllBooks() ([]models.Book, error)
}

// GetBook - retrieves books by their id from the database
func (s *Service) GetBookById(id uint) (models.Book, error) {
	var book models.Book
	if result := s.DB.First(&book, id); result.Error != nil {
		return models.Book{}, result.Error
	}
	return book, nil
}

// GetBooksByTitle - retrieves all books by title (path - /book/title/)
func (s *Service) GetBookByTitle(title string) ([]models.Book, error) {
	var books []models.Book
	if result := s.DB.Find(&books).Where("title = ?", title); result.Error != nil {
		return []models.Book{}, result.Error
	}
	return books, nil
}

// PostBook - adds a new book to the database
func (s *Service) PostBook(book models.Book) (models.Book, error) {
	if result := s.DB.Save(&book); result.Error != nil {
		return models.Book{}, result.Error
	}
	return book, nil
}

// UpdateBook - updates a book by ID with new book info
func (s *Service) UpdateBook(id uint, newBook models.Book) (models.Book, error) {
	book, err := s.GetBookById(id)
	if err != nil {
		return models.Book{}, err
	}

	if result := s.DB.Model(&book).Updates(newBook); result.Error != nil {
		return models.Book{}, result.Error
	}

	return book, nil
}

// DeleteBook - deletes a book from the database by Id
func (s *Service) DeleteBook(id uint) error {
	if result := s.DB.Delete(&models.Book{}, id); result.Error != nil {
		return result.Error
	}
	return nil
}

// GetAllBooks - retrives all books from the database
func (s *Service) GetAllBooks() ([]models.Book, error) {
	var books []models.Book
	if result := s.DB.Find(&books); result.Error != nil {
		return books, result.Error
	}
	return books, nil
}
