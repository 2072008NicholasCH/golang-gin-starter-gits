package resource

import (
	"gin-starter-gits/entity"

	"github.com/google/uuid"
)

type GetAuthorByIDRequest struct {
	UUID string `uri:"uuid" binding:"required"`
}

type CreateAuthorRequest struct {
	NAME   string `json:"name" binding:"required"`
	GENDER string `json:"gender" binding:"required"`
}

type UpdateAuthorRequest struct {
	NAME   string `json:"name" binding:"required"`
	GENDER string `json:"gender" binding:"required"`
}

type DeleteAuthorRequest struct {
	UUID string `uri:"uuid" binding:"required"`
}

type Author struct {
	UUID   uuid.UUID `json:"uuid"`
	NAME   string    `json:"name"`
	GENDER string    `json:"gender"`
}

func NewAuthorResponse(author *entity.Author) *Author {
	return &Author{
		UUID:   author.UUID,
		NAME:   author.Name,
		GENDER: author.Gender,
	}
}

type AuthorListResponse struct {
	List  []*Author `json:"list"`
	Total int64     `json:"total"`
}
