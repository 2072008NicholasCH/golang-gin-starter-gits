package handler

import (
	"gin-starter-gits/modules/book/v1/service"
	"gin-starter-gits/resource"
	"gin-starter-gits/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type BookUpdaterHandler struct {
	bookUpdater service.BookUpdaterUseCase
	bookFinder  service.BookFinderUseCase
}

// NewBookUpdaterHandler is a constructor for BookUpdaterHandler
func NewBookUpdaterHandler(bookUpdater service.BookUpdaterUseCase, bookFinder service.BookFinderUseCase) *BookUpdaterHandler {
	return &BookUpdaterHandler{bookUpdater, bookFinder}
}

// UpdateBook is a handler for updating a book
func (buh *BookUpdaterHandler) UpdateBook(c *gin.Context) {
	var req resource.UpdateBookRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorAPIResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
	}

	bookUUID, _ := uuid.Parse(c.Param("uuid"))

	book, err := buh.bookFinder.GetBookByID(c, bookUUID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorAPIResponse(http.StatusInternalServerError, err.Error()))
		c.Abort()
		return
	}

	book.Title = req.TITLE
	book.PublisherID = req.PUBLISHER_ID
	book.AuthorID = req.AUTHOR_ID
	book.ISBN = req.ISBN

	if err := buh.bookUpdater.UpdateBook(c, book); err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorAPIResponse(http.StatusInternalServerError, err.Error()))
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, response.SuccessAPIResponseList(http.StatusOK, "book updated", book))
}
