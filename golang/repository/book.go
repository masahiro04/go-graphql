package repository

import (
	"golang/graph/model"
	"golang/models"

	"gorm.io/gorm"
)

type BookRepository interface {
	CreateBook(bookInput *model.BookInput) (*models.Book, error)
	UpdateBook(bookInput *model.BookInput, id int) error
	DeleteBook(id int) error
	GetOneBook(id int) (*models.Book, error)
	GetAllBooks() ([]*model.Book, error)
}

type BookService struct {
	DB *gorm.DB
}

var _ BookRepository = &BookService{}

func NewBookService(db *gorm.DB) *BookService {
	return &BookService{
		DB: db,
	}
}

func (b *BookService) CreateBook(bookInput *model.BookInput) (*models.Book, error) {
	book := &models.Book{
		Title:     bookInput.Title,
		Author:    bookInput.Author,
		Publisher: bookInput.Publisher,
	}
	err := b.DB.Create(&book).Error
	return book, err
}

func (b *BookService) UpdateBook(bookInput *model.BookInput, id int) error {
	book := model.Book{
		ID:        id,
		Title:     bookInput.Title,
		Author:    bookInput.Author,
		Publisher: bookInput.Publisher,
	}
	err := b.DB.Model(&book).Where("id = ?", id).Updates(book).Error
	return err
}

func (b *BookService) DeleteBook(id int) error {
	book := &models.Book{}
	err := b.DB.Delete(book, id).Error
	return err
}

func (b *BookService) GetOneBook(id int) (*models.Book, error) {
	book := &models.Book{}
	err := b.DB.Where("id = ?", id).First(book).Error
	return book, err
}

func (b *BookService) GetAllBooks() ([]*model.Book, error) {
	books := []*model.Book{}
	err := b.DB.Find(&books).Error
	return books, err
}
