package resource

import (
	"gin-starter-gits/entity"

	"github.com/google/uuid"
)

type GetBookByIDRequest struct {
	UUID string `uri:"uuid" binding:"required"`
}

type CreateBookRequest struct {
	TITLE        string `json:"title" binding:"required"`
	ISBN         string `json:"isbn" binding:"required"`
	AUTHOR_ID    int64  `json:"author_id" binding:"required"`
	PUBLISHER_ID int64  `json:"publisher_id" binding:"required"`
}

type UpdateBookRequest struct {
	TITLE        string `json:"title"`
	ISBN         string `json:"isbn"`
	AUTHOR_ID    int64  `json:"author_id"`
	PUBLISHER_ID int64  `json:"publisher_id"`
}

type DeleteBookRequest struct {
	UUID string `uri:"uuid" binding:"required"`
}

type Book struct {
	UUID         uuid.UUID `json:"uuid"`
	ISBN         string    `json:"isbn"`
	TITLE        string    `json:"name"`
	AUTHOR_ID    int64     `json:"author_id"`
	PUBLISHER_ID int64     `json:"publisher_id"`
}

func NewBookResponse(book *entity.Book) *Book {
	return &Book{
		UUID:         book.UUID,
		ISBN:         book.ISBN,
		TITLE:        book.Title,
		AUTHOR_ID:    book.AuthorID,
		PUBLISHER_ID: book.PublisherID,
	}
}

type BookListResponse struct {
	List  []*Book `json:"list"`
	Total int64   `json:"total"`
}
