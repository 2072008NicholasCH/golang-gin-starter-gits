package resource

import "gin-starter-gits/entity"

type GetAuthorByIDRequest struct {
	ID int64 `uri:"id" binding:"required"`
}

type Author struct {
	ID     int64  `json:"id"`
	NAME   string `json:"name"`
	GENDER string `json:"gender"`
}

func NewAuthorResponse(author *entity.Author) *Author {
	return &Author{
		ID:     author.ID,
		NAME:   author.Name,
		GENDER: author.Gender,
	}
}

type AuthorListResponse struct {
	List  []*Author `json:"list"`
	Total int64     `json:"total"`
}
