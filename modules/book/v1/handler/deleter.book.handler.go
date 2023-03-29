package handler

import (
	"gin-starter-gits/modules/book/v1/service"
	"gin-starter-gits/resource"
	"gin-starter-gits/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type BookDeleterHandler struct {
	bookDeleter service.BookDeleterUseCase
}

// NewBookDeleterHandler is a constructor for BookDeleterHandler
func NewBookDeleterHandler(bookDeleter service.BookDeleterUseCase) *BookDeleterHandler {
	return &BookDeleterHandler{bookDeleter}
}

// DeleteBook is a handler for deleting a book
func (bdh *BookDeleterHandler) DeleteBook(c *gin.Context) {
	var req resource.DeleteBookRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorAPIResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
	}

	bookUUID, _ := uuid.Parse(req.UUID)
	if err := bdh.bookDeleter.DeleteBook(c, bookUUID); err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorAPIResponse(http.StatusInternalServerError, err.Error()))
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, response.SuccessAPIResponseList(http.StatusOK, "book deleted", nil))
}
