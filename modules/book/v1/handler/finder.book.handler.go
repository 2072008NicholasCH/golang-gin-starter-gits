package handler

import (
	"gin-starter-gits/common/errors"
	service "gin-starter-gits/modules/book/v1/service"
	"gin-starter-gits/resource"
	"gin-starter-gits/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BookFinderHandler struct {
	bookFinder service.BookFinderUseCase
}

// NewBookHandler is a constructor for BookHandler
func NewBookHandler(
	bookFinder service.BookFinderUseCase,
) *BookFinderHandler {
	return &BookFinderHandler{bookFinder}
}

// Get Books
func (bf *BookFinderHandler) GetBooks(c *gin.Context) {
	books, err := bf.bookFinder.GetBooks(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": books})
}

// Get Book by ID
func (bf *BookFinderHandler) GetBookByID(c *gin.Context) {
	var req resource.GetBookByIDRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorAPIResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
	}

	book, err := bf.bookFinder.GetBookByID(c.Request.Context(), req.ID)
	if err != nil {
		c.JSON(errors.ErrInternalServerError.Code, response.ErrorAPIResponse(errors.ErrInternalServerError.Code, err.Error()))
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, response.SuccessAPIResponseList(http.StatusOK, "success", resource.NewBookResponse(book)))

}
