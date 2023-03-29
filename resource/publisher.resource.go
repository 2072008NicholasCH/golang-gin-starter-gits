package resource

import (
	"gin-starter-gits/entity"

	"github.com/google/uuid"
)

type GetPublisherByUUIDRequest struct {
	UUID string `uri:"uuid" binding:"required"`
}

type CreatePublisherRequest struct {
	NAME string `json:"name" binding:"required"`
	KOTA string `json:"kota" binding:"required"`
}

type UpdatePublisherRequest struct {
	NAME string `json:"name" binding:"required"`
	KOTA string `json:"kota" binding:"required"`
}

type DeletePublisherRequest struct {
	UUID string `uri:"uuid" binding:"required"`
}

type Publisher struct {
	UUID uuid.UUID `json:"uuid"`
	NAME string    `json:"name"`
	KOTA string    `json:"kota"`
}

func NewPublisherResponse(publisher *entity.Publisher) *Publisher {
	return &Publisher{
		UUID: publisher.UUID,
		NAME: publisher.Name,
		KOTA: publisher.Kota,
	}
}

type PublisherListResponse struct {
	List  []*Publisher `json:"list"`
	Total int64        `json:"total"`
}
