package server

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetRouter() *gin.Engine {

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{http.MethodGet, http.MethodPatch, http.MethodPost, http.MethodHead, http.MethodDelete, http.MethodOptions},
		AllowHeaders:     []string{"Content-Type", "X-XSRF-TOKEN", "Accept", "Origin", "X-Requested-With", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	router.RedirectTrailingSlash = true
	api := router.Group("/api")
	{
		api.POST("/signup", signUp)
		api.POST("/signin", signIn)
		api.POST("/getUsers", getUsers)
		api.POST("/updateProfile", updateProfile)
		api.POST("/getProfiles", getProfiles)
		//api.POST("/testQuery", testQuery)
		//api.POST("/signUp2", signUp2)
	}

	router.NoRoute(func(ctx *gin.Context) { ctx.JSON(http.StatusNotFound, gin.H{}) })

	return router
}
