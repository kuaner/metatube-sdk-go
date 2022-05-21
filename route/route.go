package route

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/javtube/javtube-sdk-go/engine"
	"github.com/javtube/javtube-sdk-go/route/validator"
)

func New(app *engine.Engine, v validator.Validator) *gin.Engine {
	r := gin.New()
	{
		// register middleware
		r.Use(logger(), recovery())
		// fallback behavior
		r.NoRoute(notFound())
		r.NoMethod(notAllowed())
	}

	api := r.Group("/api")
	api.Use(authentication(v))
	{
		// info/metadata
		api.GET("/actor", getInfo(app, actorInfoType))
		api.GET("/movie", getInfo(app, movieInfoType))

		// search
		search := api.Group("/search")
		search.GET("/actor", getSearchResults(app, actorSearchType))
		search.GET("/movie", getSearchResults(app, movieSearchType))
	}

	img := r.Group("/image")
	{
		img.GET("/primary", getImage(app, primaryImageType))
		img.GET("/thumb", getImage(app, thumbImageType))
		img.GET("/backdrop", getImage(app, backdropImageType))
	}

	return r
}

func logger() gin.HandlerFunc {
	return gin.LoggerWithConfig(gin.LoggerConfig{})
}

func recovery() gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, err any) {
		abortWithStatusMessage(c, http.StatusInternalServerError, err)
	})
}

func notFound() gin.HandlerFunc {
	return func(c *gin.Context) {
		abortWithStatusMessage(c, http.StatusNotFound,
			http.StatusText(http.StatusNotFound))
	}
}

func notAllowed() gin.HandlerFunc {
	return func(c *gin.Context) {
		abortWithStatusMessage(c, http.StatusMethodNotAllowed,
			http.StatusText(http.StatusMethodNotAllowed))
	}
}

func abortWithStatusMessage(c *gin.Context, code int, message any) {
	switch m := message.(type) {
	case error:
		message = m.Error()
	case fmt.Stringer:
		message = m.String()
	}
	c.AbortWithStatusJSON(code, &struct {
		Status  bool `json:"status"`
		Message any  `json:"message"`
	}{
		Status:  false,
		Message: message,
	})
}
