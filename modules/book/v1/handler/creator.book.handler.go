package handler

import (
	"gin-starter-gits/modules/book/v1/service"
	"gin-starter-gits/resource"
	"gin-starter-gits/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BookCreatorHandler struct {
	bookCreator service.BookCreatorUseCase
}

// NewBookCreatorHandler is a constructor for BookCreatorHandler
func NewBookCreatorHandler(bookCreator service.BookCreatorUseCase) *BookCreatorHandler {
	return &BookCreatorHandler{bookCreator}
}

// CreateBook is a handler for creating a book
func (bch *BookCreatorHandler) CreateBook(c *gin.Context) {
	var req resource.CreateBookRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorAPIResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
	}

	book, err := bch.bookCreator.CreateBook(c, req.ISBN, req.TITLE, req.AUTHOR_ID, req.PUBLISHER_ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorAPIResponse(http.StatusInternalServerError, err.Error()))
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, response.SuccessAPIResponseList(http.StatusOK, "book created", book))
}
