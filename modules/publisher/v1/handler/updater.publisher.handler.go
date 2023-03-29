package handler

import (
	"gin-starter-gits/common/errors"
	"gin-starter-gits/modules/publisher/v1/service"
	"gin-starter-gits/resource"
	"gin-starter-gits/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PublisherUpdaterHandler struct {
	publisherUpdater service.PublisherUpdaterUseCase
	publisherFinder  service.PublisherFinderUseCase
}

// NewPublisherUpdaterHandler is a constructor for PublisherUpdaterHandler
func NewPublisherUpdaterHandler(
	publisherUpdater service.PublisherUpdaterUseCase,
	publisherFinder service.PublisherFinderUseCase,
) *PublisherUpdaterHandler {
	return &PublisherUpdaterHandler{publisherUpdater, publisherFinder}
}

// Update Publisher is a handler for updating publisher
func (puph *PublisherUpdaterHandler) UpdatePublisher(c *gin.Context) {
	var req resource.UpdatePublisherRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorAPIResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
	}

	pubUUID, _ := uuid.Parse(c.Param("uuid"))

	publisher, err := puph.publisherFinder.GetPublisherByID(c, pubUUID)
	if err != nil {
		c.JSON(errors.ErrInternalServerError.Code, response.ErrorAPIResponse(errors.ErrInternalServerError.Code, err.Error()))
		c.Abort()
		return
	}

	publisher.Name = req.NAME
	publisher.Kota = req.KOTA

	if err := puph.publisherUpdater.UpdatePublisher(c, publisher); err != nil {
		c.JSON(errors.ErrInternalServerError.Code, response.ErrorAPIResponse(errors.ErrInternalServerError.Code, err.Error()))
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, response.SuccessAPIResponseList(http.StatusOK, "success", resource.NewPublisherResponse(publisher)))
}
