package handler

import (
	"gin-starter-gits/common/errors"
	"gin-starter-gits/modules/author/v1/service"
	"gin-starter-gits/resource"
	"gin-starter-gits/response"
	"net/http"

	"github.com/gin-gonic/gin"
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
	c.JSON(http.StatusOK, gin.H{"data": authors})
}

// Get Author by ID
func (af *AuthorFinderHandler) GetAuthorByID(c *gin.Context) {
	var req resource.GetAuthorByIDRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorAPIResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
	}

	author, err := af.authorFinder.GetAuthorByID(c.Request.Context(), req.ID)
	if err != nil {
		c.JSON(errors.ErrInternalServerError.Code, response.ErrorAPIResponse(errors.ErrInternalServerError.Code, err.Error()))
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, response.SuccessAPIResponseList(http.StatusOK, "success", resource.NewAuthorResponse(author)))

}
