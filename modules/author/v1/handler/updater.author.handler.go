package handler

import (
	"gin-starter-gits/modules/author/v1/service"
	"gin-starter-gits/resource"
	"gin-starter-gits/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AuthorUpdaterHandler struct {
	authorUpdater service.AuthorUpdaterUseCase
	authorFinder  service.AuthorFinderUseCase
}

// NewAuthorUpdaterHandler is a function to initialize author updater handler
func NewAuthorUpdaterHandler(authorUpdater service.AuthorUpdaterUseCase, authorFinder service.AuthorFinderUseCase) *AuthorUpdaterHandler {
	return &AuthorUpdaterHandler{
		authorUpdater: authorUpdater,
		authorFinder:  authorFinder,
	}
}

// UpdateAuthor is a function to update author
func (auh *AuthorUpdaterHandler) UpdateAuthor(c *gin.Context) {
	var req resource.UpdateAuthorRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorAPIResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
	}

	authUUID, _ := uuid.Parse(c.Param("uuid"))
	author, err := auh.authorFinder.GetAuthorByID(c, authUUID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorAPIResponse(http.StatusInternalServerError, err.Error()))
		c.Abort()
		return
	}

	author.Name = req.NAME
	author.Gender = req.GENDER

	if err := auh.authorUpdater.UpdateAuthor(c, author); err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorAPIResponse(http.StatusInternalServerError, err.Error()))
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, response.SuccessAPIResponseList(http.StatusOK, "success", resource.NewAuthorResponse(author)))
}
