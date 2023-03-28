package resource

import "gin-starter-gits/entity"

type GetPublisherByIDRequest struct {
	ID int64 `uri:"id" binding:"required"`
}

type Publisher struct {
	ID   int64  `json:"id"`
	NAME string `json:"name"`
	KOTA string `json:"kota"`
}

func NewPublisherResponse(publisher *entity.Publisher) *Publisher {
	return &Publisher{
		ID:   publisher.ID,
		NAME: publisher.Name,
		KOTA: publisher.Kota,
	}
}

type PublisherListResponse struct {
	List  []*Publisher `json:"list"`
	Total int64        `json:"total"`
}
