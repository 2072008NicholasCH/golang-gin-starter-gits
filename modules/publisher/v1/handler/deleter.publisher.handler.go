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

type PublisherDeleterHandler struct {
	publisherDeleter service.PublisherDeleterUseCase
}

func NewPublisherDeleterHandler(publisherDeleter service.PublisherDeleterUseCase) *PublisherDeleterHandler {
	return &PublisherDeleterHandler{publisherDeleter}
}

func (pdh *PublisherDeleterHandler) DeletePublisher(c *gin.Context) {
	var req resource.DeletePublisherRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorAPIResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
	}

	pubUUID, _ := uuid.Parse(req.UUID)
	if err := pdh.publisherDeleter.DeletePublisher(c, pubUUID); err != nil {
		c.JSON(errors.ErrInternalServerError.Code, response.ErrorAPIResponse(errors.ErrInternalServerError.Code, err.Error()))
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, response.SuccessAPIResponseList(http.StatusOK, "success", nil))
}
