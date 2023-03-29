package handler

import (
	"gin-starter-gits/common/errors"
	"gin-starter-gits/modules/publisher/v1/service"
	"gin-starter-gits/resource"
	"gin-starter-gits/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PublisherCreatorHandler struct {
	publisherCreator service.PublisherCreatorUseCase
}

// NewPublisherCreatorHandler is a constructor for PublisherCreatorHandler
func NewPublisherCreatorHandler(
	publisherCreator service.PublisherCreatorUseCase,
) *PublisherCreatorHandler {
	return &PublisherCreatorHandler{publisherCreator}
}

// CreatePublisher is a handler for creating publisher
func (pch *PublisherCreatorHandler) CreatePublisher(c *gin.Context) {
	var req resource.CreatePublisherRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorAPIResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
	}

	publisher, err := pch.publisherCreator.CreatePublisher(c, req.NAME, req.KOTA)
	if err != nil {
		c.JSON(errors.ErrInternalServerError.Code, response.ErrorAPIResponse(errors.ErrInternalServerError.Code, err.Error()))
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, response.SuccessAPIResponseList(http.StatusOK, "success", resource.NewPublisherResponse(publisher)))
}
