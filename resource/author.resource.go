package resource

import (
	"gin-starter-gits/entity"

	"github.com/google/uuid"
)

type GetAuthorByIDRequest struct {
	ID int64 `uri:"id" binding:"required"`
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
