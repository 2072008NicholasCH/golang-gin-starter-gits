package handler

import (
	"gin-starter-gits/common/errors"
	service "gin-starter-gits/modules/book/v1/service"
	"gin-starter-gits/resource"
	"gin-starter-gits/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
		c.JSON(http.StatusInternalServerError, response.ErrorAPIResponse(http.StatusInternalServerError, err.Error()))
		c.Abort()
		return
	}
	res := make([]*resource.Book, 0)

	for _, book := range books {
		res = append(res, resource.NewBookResponse(book))
	}

	c.JSON(http.StatusOK, response.SuccessAPIResponseList(http.StatusOK, "success", &resource.BookListResponse{List: res, Total: int64(len(res))}))
}

// Get Book by ID
func (bf *BookFinderHandler) GetBookByID(c *gin.Context) {
	var req resource.GetBookByIDRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorAPIResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
	}

	bookUUID, _ := uuid.Parse(req.UUID)
	book, err := bf.bookFinder.GetBookByID(c.Request.Context(), bookUUID)
	if err != nil {
		c.JSON(errors.ErrInternalServerError.Code, response.ErrorAPIResponse(errors.ErrInternalServerError.Code, err.Error()))
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, response.SuccessAPIResponseList(http.StatusOK, "success", resource.NewBookResponse(book)))

}
