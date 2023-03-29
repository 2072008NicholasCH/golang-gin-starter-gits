package entity

import (
	"time"

	"github.com/google/uuid"
)

const (
	bookTableName = "main.books"
)

type Book struct {
	ID         int64     `json:"id"`
	UUID       uuid.UUID `json:"uuid"`
	Title      string    `json:"title"`
	ISBN       string    `json:"isbn"`
	AuthorID   int64     `json:"author_id"`
	PubliserID int64     `json:"publisher_id"`
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
		UUID:       uuid,
		Title:      title,
		ISBN:       isbn,
		AuthorID:   authorID,
		PubliserID: publisherID,
		Auditable:  NewAuditable(createdBy),
	}
}

func (model *Book) MapUpdateFrom(from *Book) *map[string]interface{} {
	if from == nil {
		return &map[string]interface{}{
			"title":        model.Title,
			"isbn":         model.ISBN,
			"author_id":    model.AuthorID,
			"publisher_id": model.PubliserID,
			"updated_at":   model.UpdatedAt,
		}
	}
	mapped := make(map[string]interface{})

	if model.Title != from.Title {
		mapped["title"] = from.Title
	}

	if model.ISBN != from.ISBN {
		mapped["isbn"] = from.ISBN
	}

	if model.AuthorID != from.AuthorID {
		mapped["author_id"] = from.AuthorID
	}

	if model.PubliserID != from.PubliserID {
		mapped["publisher_id"] = from.PubliserID
	}

	mapped["updated_at"] = time.Now()
	return &mapped
}
