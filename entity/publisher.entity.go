package entity

const (
	publisherTableName = "main.publishers"
)

type Publisher struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Kota string `json:"kota"`
	// Auditable
}

func (model *Publisher) TableName() string {
	return publisherTableName
}

func NewPublisher(
	name string,
	kota string,
	// createdBy string,
) *Publisher {
	return &Publisher{
		Name: name,
		Kota: kota,
		// Auditable: NewAuditable(createdBy),
	}
}

func (model *Publisher) MapUpdateFrom(from *Publisher) *map[string]interface{} {
	if from == nil {
		return &map[string]interface{}{
			"name": model.Name,
			// "updated_at": model.UpdatedAt,
		}
	}
	mapped := make(map[string]interface{})

	if model.Name != from.Name {
		mapped["name"] = from.Name
	}

	if model.Kota != from.Kota {
		mapped["kota"] = from.Kota
	}

	return &mapped
}
