package app

import (
	"gin-starter-gits/config"
	"gin-starter-gits/middleware"
	authhandlerv1 "gin-starter-gits/modules/auth/v1/handler"
	authservicev1 "gin-starter-gits/modules/auth/v1/service"
	authorhandlerv1 "gin-starter-gits/modules/author/v1/handler"
	authorservicev1 "gin-starter-gits/modules/author/v1/service"
	"gin-starter-gits/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

// DeprecatedAPI is a handler for deprecated APIs
func DeprecatedAPI(c *gin.Context) {
	c.JSON(http.StatusForbidden, response.ErrorAPIResponse(http.StatusForbidden, "this version of api is deprecated. please use another version."))
	c.Abort()
}

// DefaultHTTPHandler is a handler for default APIs
func DefaultHTTPHandler(cfg config.Config, router *gin.Engine) {
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, response.ErrorAPIResponse(http.StatusNotFound, "invalid route"))
		c.Abort()
	})
}

// AuthHTTPHandler is a handler for auth APIs
func AuthHTTPHandler(cfg config.Config, router *gin.Engine, auc authservicev1.AuthUseCase) {
	hnd := authhandlerv1.NewAuthHandler(auc)
	v1 := router.Group("/v1")
	{
		v1.POST("/user/login", hnd.Login)
	}
}

// AuthHTTPHandler is a handler for author APIs
func AuthorFinderHTTPHandler(cfg config.Config, router *gin.Engine, atc authorservicev1.AuthorFinderUseCase) {
	hnd := authorhandlerv1.NewAuthorHandler(atc)
	v1 := router.Group("/v1")

	v1.Use(middleware.Auth(cfg))

	{
		v1.GET("/author", hnd.GetAuthors)
		v1.GET("/author/detail/:id", hnd.GetAuthorByID)
	}
}
