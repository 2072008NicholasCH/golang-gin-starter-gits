package resource

import "gin-starter-gits/entity"

type GetBookByIDRequest struct {
	ID int64 `uri:"id" binding:"required"`
}

type Book struct {
	ID           int64  `json:"id"`
	TITLE        string `json:"name"`
	AUTHOR_ID    int64  `json:"author_id"`
	PUBLISHER_ID int64  `json:"publisher_id"`
}

func NewBookResponse(book *entity.Book) *Book {
	return &Book{
		ID:           book.ID,
		TITLE:        book.Title,
		AUTHOR_ID:    book.AuthorID,
		PUBLISHER_ID: book.PubliserID,
	}
}

type BookListResponse struct {
	List  []*Book `json:"list"`
	Total int64   `json:"total"`
}
