package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"golang/graph/generated"
	"golang/graph/model"
)

func (r *mutationResolver) CreateBook(ctx context.Context, input model.BookInput) (*model.Book, error) {
	book, err := r.BookRepository.CreateBook(&input)
	bookCreate := &model.Book{
		ID:        book.ID,
		Title:     book.Title,
		Author:    book.Author,
		Publisher: book.Publisher,
	}
	if err != nil {
		return nil, err
	}
	return bookCreate, nil
}

func (r *mutationResolver) DeleteBook(ctx context.Context, id int) (string, error) {
	err := r.BookRepository.DeleteBook(id)
	if err != nil {
		return "", err

	}
	successMessage := "successfully deleted"
	return successMessage, nil
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateBook(ctx context.Context, id int, input *model.BookInput) (string, error) {
	err := r.BookRepository.UpdateBook(input, id)
	if err != nil {
		return "nil", err
	}
	successMessage := "successfully updated"
	return successMessage, nil
}

func (r *queryResolver) GetAllBooks(ctx context.Context) ([]*model.Book, error) {
	books, err := r.BookRepository.GetAllBooks()
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (r *queryResolver) GetOneBook(ctx context.Context, id int) (*model.Book, error) {
	book, err := r.BookRepository.GetOneBook(id)
	selectedBook := &model.Book{
		ID:        book.ID,
		Title:     book.Title,
		Author:    book.Author,
		Publisher: book.Publisher,
	}
	if err != nil {
		return nil, err
	}
	return selectedBook, err
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
