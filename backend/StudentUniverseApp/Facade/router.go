package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setRouter() *gin.Engine {

	router := gin.Default()
	router.RedirectTrailingSlash = true
	api := router.Group("/api")
	{
		api.POST("/signup", signUp)
		api.OPTIONS("/signup", preflight)
		api.POST("/signin", signIn)
		api.POST("/getUsers", getUsers)
	}

	router.NoRoute(func(ctx *gin.Context) { ctx.JSON(http.StatusNotFound, gin.H{}) })

	return router
}

func preflight(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
	c.JSON(http.StatusOK, struct{}{})
}
