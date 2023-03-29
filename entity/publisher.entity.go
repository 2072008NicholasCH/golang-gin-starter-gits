package entity

import "github.com/google/uuid"

const (
	publisherTableName = "main.publishers"
)

type Publisher struct {
	ID   int64     `json:"id"`
	UUID uuid.UUID `json:"uuid"`
	Name string    `json:"name"`
	Kota string    `json:"kota"`
	Auditable
}

func (model *Publisher) TableName() string {
	return publisherTableName
}

func NewPublisher(
	uuid uuid.UUID,
	name string,
	kota string,
	createdBy string,
) *Publisher {
	return &Publisher{
		UUID:      uuid,
		Name:      name,
		Kota:      kota,
		Auditable: NewAuditable(createdBy),
	}
}

func (model *Publisher) MapUpdateFrom(from *Publisher) *map[string]interface{} {
	if from == nil {
		return &map[string]interface{}{
			"name":       model.Name,
			"kota":       model.Kota,
			"updated_at": model.UpdatedAt,
		}
	}
	mapped := make(map[string]interface{})

	if (model.Name != from.Name) && (from.Name != "") {
		mapped["name"] = from.Name
	}

	if (model.Kota != from.Kota) && (from.Kota != "") {
		mapped["kota"] = from.Kota
	}

	mapped["updated_at"] = model.UpdatedAt
	return &mapped
}
