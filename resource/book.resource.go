package resource

import (
	"gin-starter-gits/entity"

	"github.com/google/uuid"
)

type GetBookByIDRequest struct {
	UUID string `uri:"uuid" binding:"required"`
}

type Book struct {
	UUID         uuid.UUID `json:"uuid"`
	TITLE        string    `json:"name"`
	AUTHOR_ID    int64     `json:"author_id"`
	PUBLISHER_ID int64     `json:"publisher_id"`
}

func NewBookResponse(book *entity.Book) *Book {
	return &Book{
		UUID:         book.UUID,
		TITLE:        book.Title,
		AUTHOR_ID:    book.AuthorID,
		PUBLISHER_ID: book.PublisherID,
	}
}

type BookListResponse struct {
	List  []*Book `json:"list"`
	Total int64   `json:"total"`
}
