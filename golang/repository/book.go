package repository

import (
	"golang/graph/model"

	"gorm.io/gorm"
)

type BookRepository interface {
	CreateBook(bookInput *model.BookInput) (*model.Book, error)
	UpdateBook(bookInput *model.BookInput, id int) error
	DeleteBook(id int) error
	GetOneBook(id int) (*model.Book, error)
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

func (b *BookService) CreateBook(bookInput *model.BookInput) (*model.Book, error) {
	book := &model.Book{
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
