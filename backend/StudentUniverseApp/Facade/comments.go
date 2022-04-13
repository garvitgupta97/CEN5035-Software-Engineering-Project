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

func createComment(ctx *gin.Context) {
	comment := new(database.Comment)

	if err := ctx.Bind(comment); err != nil {

		var verr validator.ValidationErrors
		if errors.As(err, &verr) {
			ctx.JSON(http.StatusBadRequest, gin.H{"errors": SimpleErrorMsg(verr)})
			return
		}

		log.Info().Err(err).Msg("unable to bind")

		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Comment Failed": err.Error()})
		return
	}
	if !database.CreateComment(*comment) {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Comment Failed": "Error"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Comment added successfully"})
}

func getCommentsByPosts(ctx *gin.Context) {
	comment := new(database.Comment)

	if err := ctx.Bind(comment); err != nil {
		var verr validator.ValidationErrors
		if errors.As(err, &verr) {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": SimpleErrorMsg(verr)})
			return
		}
		log.Info().Err(err).Msg("unable to find comment")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"unable to find comment": err.Error()})
		return
	}
	allComments := database.GetCommentsByPosts(comment.PostId)

	ctx.JSON(http.StatusOK, allComments)
}