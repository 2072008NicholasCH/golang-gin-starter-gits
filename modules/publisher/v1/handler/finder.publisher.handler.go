package handler

import (
	"gin-starter-gits/common/errors"
	"gin-starter-gits/modules/publisher/v1/service"
	"gin-starter-gits/resource"
	"gin-starter-gits/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PublisherFinderHandler struct {
	publisherFinder service.PublisherFinderUseCase
}

// NewPublisherHandler is a constructor for PublisherHandler
func NewPublisherHandler(
	publisherFinder service.PublisherFinderUseCase,
) *PublisherFinderHandler {
	return &PublisherFinderHandler{publisherFinder}
}

// Get Publishers
func (pf *PublisherFinderHandler) GetPublishers(c *gin.Context) {
	publishers, err := pf.publisherFinder.GetPublishers(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": publishers})
}

// Get Publisher by ID
func (pf *PublisherFinderHandler) GetPublisherByID(c *gin.Context) {
	var req resource.GetPublisherByIDRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorAPIResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
	}

	publisher, err := pf.publisherFinder.GetPublisherByID(c.Request.Context(), req.ID)
	if err != nil {
		c.JSON(errors.ErrInternalServerError.Code, response.ErrorAPIResponse(errors.ErrInternalServerError.Code, err.Error()))
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, response.SuccessAPIResponseList(http.StatusOK, "success", resource.NewPublisherResponse(publisher)))

}
