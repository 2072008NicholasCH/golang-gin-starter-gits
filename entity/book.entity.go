package entity

import (
	"time"

	"github.com/google/uuid"
)

const (
	bookTableName = "main.books"
)

type Book struct {
	ID          int64     `json:"id"`
	UUID        uuid.UUID `json:"uuid"`
	Title       string    `json:"title"`
	ISBN        string    `json:"isbn"`
	AuthorID    int64     `json:"author_id"`
	PublisherID int64     `json:"publisher_id"`
	Auditable
}

func (model *Book) TableName() string {

	return bookTableName
}

func NewBook(
	uuid uuid.UUID,
	title string,
	isbn string,
	authorID int64,
	publisherID int64,
	createdBy string,
) *Book {
	return &Book{
		UUID:        uuid,
		Title:       title,
		ISBN:        isbn,
		AuthorID:    authorID,
		PublisherID: publisherID,
		Auditable:   NewAuditable(createdBy),
	}
}

func (model *Book) MapUpdateFrom(from *Book) *map[string]interface{} {
	if from == nil {
		return &map[string]interface{}{
			"title":        model.Title,
			"isbn":         model.ISBN,
			"author_id":    model.AuthorID,
			"publisher_id": model.PublisherID,
			"updated_at":   model.UpdatedAt,
		}
	}
	mapped := make(map[string]interface{})

	if (model.Title != from.Title) && (from.Title != "") {
		mapped["title"] = from.Title
	}

	if (model.ISBN != from.ISBN) && (from.ISBN != "") {
		mapped["isbn"] = from.ISBN
	}

	if (model.AuthorID != from.AuthorID) && (from.AuthorID != 0) {
		mapped["author_id"] = from.AuthorID
	}

	if (model.PublisherID != from.PublisherID) && (from.PublisherID != 0) {
		mapped["publisher_id"] = from.PublisherID
	}

	mapped["updated_at"] = time.Now()
	return &mapped
}
