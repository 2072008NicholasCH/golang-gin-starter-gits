package handler

import (
	"gin-starter-gits/modules/author/v1/service"
	"gin-starter-gits/resource"
	"gin-starter-gits/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AuthorDeleterHandler struct {
	authorDeleter service.AuthorDeleterUseCase
}

// NewAuthorDeleterHandler is a function to initialize author deleter handler
func NewAuthorDeleterHandler(authorDeleter service.AuthorDeleterUseCase) *AuthorDeleterHandler {
	return &AuthorDeleterHandler{
		authorDeleter: authorDeleter,
	}
}

// DeleteAuthor is a function to delete author
func (adh *AuthorDeleterHandler) DeleteAuthor(c *gin.Context) {
	var req resource.DeleteAuthorRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorAPIResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
	}

	autUUID, _ := uuid.Parse(req.UUID)
	if err := adh.authorDeleter.DeleteAuthor(c, autUUID); err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorAPIResponse(http.StatusInternalServerError, err.Error()))
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, response.SuccessAPIResponseList(http.StatusOK, "success", nil))
}
