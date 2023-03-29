package handler

import (
	"gin-starter-gits/common/errors"
	"gin-starter-gits/modules/author/v1/service"
	"gin-starter-gits/resource"
	"gin-starter-gits/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AuthorFinderHandler struct {
	authorFinder service.AuthorFinderUseCase
}

// NewAuthorHandler is a constructor for AuthorHandler
func NewAuthorHandler(
	authorFinder service.AuthorFinderUseCase,
) *AuthorFinderHandler {
	return &AuthorFinderHandler{authorFinder}
}

// Get Authors
func (af *AuthorFinderHandler) GetAuthors(c *gin.Context) {
	authors, err := af.authorFinder.GetAuthors(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	res := make([]*resource.Author, 0)

	for _, author := range authors {
		res = append(res, resource.NewAuthorResponse(author))
	}

	c.JSON(http.StatusOK, response.SuccessAPIResponseList(http.StatusOK, "success", &resource.AuthorListResponse{List: res, Total: int64(len(res))}))
}

// Get Author by ID
func (af *AuthorFinderHandler) GetAuthorByID(c *gin.Context) {
	var req resource.GetAuthorByIDRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorAPIResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
	}

	autUUID, _ := uuid.Parse(req.UUID)
	author, err := af.authorFinder.GetAuthorByID(c.Request.Context(), autUUID)
	if err != nil {
		c.JSON(errors.ErrInternalServerError.Code, response.ErrorAPIResponse(errors.ErrInternalServerError.Code, err.Error()))
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, response.SuccessAPIResponseList(http.StatusOK, "success", resource.NewAuthorResponse(author)))

}
