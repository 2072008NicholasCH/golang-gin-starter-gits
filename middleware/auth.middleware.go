package middleware

import (
	"gin-starter-gits/config"
	"gin-starter-gits/response"
	"gin-starter-gits/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var UserID uuid.UUID

func Auth(cfg config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := strings.Split(c.Request.Header.Get("Authorization"), "Bearer ")

		if len(tokenString) < 2 {
			c.JSON(http.StatusUnauthorized, response.ErrorAPIResponse(http.StatusUnauthorized, "unauthorized"))
			c.Abort()
			return
		}

		claims, err := utils.JWTDecode(cfg, tokenString[1])

		if err != nil {
			c.JSON(http.StatusUnauthorized, response.ErrorAPIResponse(http.StatusUnauthorized, err.Error()))
			c.Abort()
			return
		}

		UserID = claims.Subject

		c.Next()
	}
}
