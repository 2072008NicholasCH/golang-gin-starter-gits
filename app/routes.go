package app

import (
	"gin-starter-gits/config"
	"gin-starter-gits/middleware"
	authhandlerv1 "gin-starter-gits/modules/auth/v1/handler"
	authservicev1 "gin-starter-gits/modules/auth/v1/service"
	authorhandlerv1 "gin-starter-gits/modules/author/v1/handler"
	authorservicev1 "gin-starter-gits/modules/author/v1/service"
	bookhandlerv1 "gin-starter-gits/modules/book/v1/handler"
	bookservicev1 "gin-starter-gits/modules/book/v1/service"
	publisherhandlerv1 "gin-starter-gits/modules/publisher/v1/handler"
	publisherservicev1 "gin-starter-gits/modules/publisher/v1/service"
	"gin-starter-gits/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

//==================================================================================================

// DeprecatedAPI is a handler for deprecated APIs
func DeprecatedAPI(c *gin.Context) {
	c.JSON(
		http.StatusForbidden,
		response.ErrorAPIResponse(
			http.StatusForbidden, "this version of api is deprecated. please use another version.",
		),
	)
	c.Abort()
}

// DefaultHTTPHandler is a handler for default APIs
func DefaultHTTPHandler(cfg config.Config, router *gin.Engine) {
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, response.ErrorAPIResponse(http.StatusNotFound, "invalid route"))
		c.Abort()
	})
}

//==================================================================================================

// AuthHTTPHandler is a handler for auth APIs
func AuthHTTPHandler(cfg config.Config, router *gin.Engine, auc authservicev1.AuthUseCase) {
	hnd := authhandlerv1.NewAuthHandler(auc)
	v1 := router.Group("/v1")
	{
		v1.POST("/user/login", hnd.Login)
	}
}

//==================================================================================================

// AuthHTTPHandler is a handler for author APIs
func AuthorFinderHTTPHandler(cfg config.Config, router *gin.Engine, atfuc authorservicev1.AuthorFinderUseCase) {
	hnd := authorhandlerv1.NewAuthorHandler(atfuc)
	v1 := router.Group("/v1")

	v1.Use(middleware.Auth(cfg))

	{
		v1.GET("/author", hnd.GetAuthors)
		v1.GET("/author/detail/:uuid", hnd.GetAuthorByID)
	}
}

//==================================================================================================

// PublisherHTTPHandler is a handler for publisher APIs
func PublisherFinderHTTPHandler(
	cfg config.Config,
	router *gin.Engine,
	pfuc publisherservicev1.PublisherFinderUseCase,
) {
	hnd := publisherhandlerv1.NewPublisherHandler(pfuc)
	v1 := router.Group("/v1")

	v1.Use(middleware.Auth(cfg))

	{
		v1.GET("/publisher", hnd.GetPublishers)
		v1.GET("/publisher/detail/:uuid", hnd.GetPublisherByID)
	}
}

// PubliserCreatorHTTPHandler is a handler for publisher APIs
func PublisherCreatorHTTPHandler(
	cfg config.Config,
	router *gin.Engine,
	pfuc publisherservicev1.PublisherCreatorUseCase,
) {
	hnd := publisherhandlerv1.NewPublisherCreatorHandler(pfuc)
	v1 := router.Group("/v1")

	v1.Use(middleware.Auth(cfg))

	{
		v1.POST("/publisher", hnd.CreatePublisher)
	}
}

// PubliserUpdaterHTTPHandler is a handler for publisher APIs
func PublisherUpdaterHTTPHandler(
	cfg config.Config,
	router *gin.Engine,
	puuc publisherservicev1.PublisherUpdaterUseCase,
	pfuc publisherservicev1.PublisherFinderUseCase,
) {
	hnd := publisherhandlerv1.NewPublisherUpdaterHandler(puuc, pfuc)
	v1 := router.Group("/v1")

	v1.Use(middleware.Auth(cfg))

	{
		v1.PUT("/publisher/:uuid", hnd.UpdatePublisher)
	}
}

// PubliserDeleterHTTPHandler is a handler for publisher APIs
func PublisherDeleterHTTPHandler(
	cfg config.Config,
	router *gin.Engine,
	pduc publisherservicev1.PublisherDeleterUseCase,
) {
	hnd := publisherhandlerv1.NewPublisherDeleterHandler(pduc)
	v1 := router.Group("/v1")

	v1.Use(middleware.Auth(cfg))

	{
		v1.DELETE("/publisher/:uuid", hnd.DeletePublisher)
	}
}

//==================================================================================================

// BookHTTPHandler is a handler for book APIs
func BookFinderHTTPHandler(cfg config.Config, router *gin.Engine, btc bookservicev1.BookFinderUseCase) {
	hnd := bookhandlerv1.NewBookHandler(btc)
	v1 := router.Group("/v1")

	v1.Use(middleware.Auth(cfg))

	{
		v1.GET("/book", hnd.GetBooks)
		v1.GET("/book/detail/:uuid", hnd.GetBookByID)
	}
}

//==================================================================================================
