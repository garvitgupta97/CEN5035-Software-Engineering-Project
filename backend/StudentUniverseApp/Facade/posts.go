package server

import (
	//store "StudentUniverse/StudentUniverseApp/Facade/DTO"
	"errors"

	database "StudentUniverse/StudentUniverseApp/Facade/Database"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	log "github.com/rs/zerolog/log"
)

func createPost(ctx *gin.Context) {
	post := new(database.Post)
	
	if err := ctx.Bind(post); err != nil {

		var verr validator.ValidationErrors
		if errors.As(err, &verr) {
			ctx.JSON(http.StatusBadRequest, gin.H{"errors": SimpleErrorMsg(verr)})
			return
		}

		log.Info().Err(err).Msg("unable to bind")

		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Post Fail": err.Error()})
		return
	}
	if !database.CreatePost(*post) {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Post Fail": "Error"})
		return
	}

}

func getPostById(ctx *gin.Context) {
	post := new(database.Post)

	if err := ctx.Bind(post); err != nil {
		var verr validator.ValidationErrors
		if errors.As(err, &verr) {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": SimpleErrorMsg(verr)})
			return
		}
		log.Info().Err(err).Msg("unable to find")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"unable to find post": err.Error()})
		return
	}
	fullPost := database.GetPostById(post.PostId)
	if string(fullPost.UserId) == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": "Post not found"})
		return
	}

	ctx.JSON(http.StatusOK, fullPost)
}
