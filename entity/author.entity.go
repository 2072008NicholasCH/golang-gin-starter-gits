package entity

import (
	"time"

	"github.com/google/uuid"
)

const (
	authorTableName = "main.authors"
)

type Author struct {
	ID     int64     `json:"id"`
	UUID   uuid.UUID `json:"uuid"`
	Name   string    `json:"name"`
	Gender string    `json:"gender"`
	Auditable
}

func (model *Author) TableName() string {
	return authorTableName
}

func NewAuthor(
	uuid uuid.UUID,
	name string,
	gender string,
	createdBy string,
) *Author {
	return &Author{
		UUID:      uuid,
		Name:      name,
		Gender:    gender,
		Auditable: NewAuditable(createdBy),
	}
}

func (model *Author) MapUpdateFrom(from *Author) *map[string]interface{} {
	if from == nil {
		return &map[string]interface{}{
			"name":       model.Name,
			"gender":     model.Gender,
			"updated_at": model.UpdatedAt,
		}
	}
	mapped := make(map[string]interface{})

	if (model.Name != from.Name) && (from.Name != "") {
		mapped["name"] = from.Name
	}

	if (model.Gender != from.Gender) && (from.Gender != "") {
		mapped["gender"] = from.Gender
	}

	mapped["updated_at"] = time.Now()
	return &mapped
}
