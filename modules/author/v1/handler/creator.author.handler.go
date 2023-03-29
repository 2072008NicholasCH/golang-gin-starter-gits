package handler

import (
	"gin-starter-gits/modules/author/v1/service"
	"gin-starter-gits/resource"
	"gin-starter-gits/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthorCreatorHandler struct {
	authorCreator service.AuthorCreatorUseCase
}

// NewAuthorCreatorHandler is a function to initialize author creator handler
func NewAuthorCreatorHandler(authorCreator service.AuthorCreatorUseCase) *AuthorCreatorHandler {
	return &AuthorCreatorHandler{
		authorCreator: authorCreator,
	}
}

// CreateAuthor is a function to create new author
func (ach *AuthorCreatorHandler) CreateAuthor(c *gin.Context) {
	var req resource.CreateAuthorRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorAPIResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return

	}

	author, err := ach.authorCreator.CreateAuthor(c, req.NAME, req.GENDER)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorAPIResponse(http.StatusInternalServerError, err.Error()))
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, response.SuccessAPIResponseList(http.StatusOK, "success", resource.NewAuthorResponse(author)))

}
