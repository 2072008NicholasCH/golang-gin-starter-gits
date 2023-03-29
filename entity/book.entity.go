package entity

const (
	bookTableName = "main.books"
)

type Book struct {
	ID         int64  `json:"id"`
	Title      string `json:"title"`
	AuthorID   int64  `json:"author_id"`
	PubliserID int64  `json:"publisher_id"`
	Auditable
}

func (model *Book) TableName() string {

	return bookTableName
}

func NewBook(
	title string,
	authorID int64,
	publisherID int64,
	createdBy string,
) *Book {
	return &Book{
		Title:      title,
		AuthorID:   authorID,
		PubliserID: publisherID,
		Auditable:  NewAuditable(createdBy),
	}
}

func (model *Book) MapUpdateFrom(from *Book) *map[string]interface{} {
	if from == nil {
		return &map[string]interface{}{
			"title":        model.Title,
			"author_id":    model.AuthorID,
			"publisher_id": model.PubliserID,
			"updated_at":   model.UpdatedAt,
		}
	}
	mapped := make(map[string]interface{})

	if model.Title != from.Title {
		mapped["title"] = from.Title
	}

	if model.AuthorID != from.AuthorID {
		mapped["author_id"] = from.AuthorID
	}

	if model.PubliserID != from.PubliserID {
		mapped["publisher_id"] = from.PubliserID
	}

	return &mapped
}
