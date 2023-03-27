package entity

const (
	authorTableName = "main.authors"
)

type Author struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Auditable
}

func (model *Author) TableName() string {
	return authorTableName
}

func NewAuthor(
	name string,
	createdBy string,
) *Author {
	return &Author{
		Name:      name,
		Auditable: NewAuditable(createdBy),
	}
}

func (model *Author) MapUpdateFrom(from *Author) *map[string]interface{} {
	if from == nil {
		return &map[string]interface{}{
			"name":       model.Name,
			"updated_at": model.UpdatedAt,
		}
	}
	mapped := make(map[string]interface{})

	if model.Name != from.Name {
		mapped["name"] = from.Name
	}

	return &mapped
}
