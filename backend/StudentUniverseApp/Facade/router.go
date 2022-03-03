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
		api.POST("/signin", signIn)
		api.POST("/getUsers", getUsers)
		api.POST("/updateProfile", updateProfile)
		api.POST("/getProfiles", getProfiles)
		api.POST("/createPost", createPost)
		api.POST("/getPosts", getPosts)
		//api.POST("/testQuery", testQuery)
		//api.POST("/signUp2", signUp2)
	}

	router.NoRoute(func(ctx *gin.Context) { ctx.JSON(http.StatusNotFound, gin.H{}) })

	return router
}
