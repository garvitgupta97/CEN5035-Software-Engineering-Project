package server

import (
	store "StudentUniverse/StudentUniverseApp/Facade/DTO"
	"errors"
	"fmt"

	database "StudentUniverse/StudentUniverseApp/Facade/Database"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"

	log "github.com/rs/zerolog/log"
)

func signUp(ctx *gin.Context) {
	user := new(store.User)
	for _, u := range store.Users {
		if u.Email == user.Email {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"err": "User already exists."})
		}
	}
	if err := ctx.Bind(user); err != nil {

		var verr validator.ValidationErrors
		if errors.As(err, &verr) {
			ctx.JSON(http.StatusBadRequest, gin.H{"errors": Simple(verr)})
			return
		}

		log.Info().Err(err).Msg("unable to bind")

		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Signup Fail": err.Error()})
		return
	}
	database.InsertStudent(user.Email, user.Password)
	store.Users = append(store.Users, user)

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Signed up successfully.",
		"jwt": "123456789",
	})
}

func Simple(verr validator.ValidationErrors) map[string]string {
	errs := make(map[string]string)

	for _, f := range verr {
		err := f.ActualTag()
		if f.Param() != "" {
			err = fmt.Sprintf("%s=%s", err, f.Param())
		}
		errs[f.Field()] = err
	}

	return errs
}

func signIn(ctx *gin.Context) {
	user := new(store.User)
	if err := ctx.Bind(user); err != nil {

		var verr validator.ValidationErrors
		if errors.As(err, &verr) {
			ctx.JSON(http.StatusBadRequest, gin.H{"errors": Simple(verr)})
			return
		}

		log.Info().Err(err).Msg("unable to bind")

		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Sign in failed": err.Error()})
		return

	}
	for _, u := range store.Users {
		if u.Email == user.Email && u.Password == user.Password {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "Signed in successfully.",
				"jwt": "123456789",
			})
			return
		}
	}
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Sign in failed": "Sign in failed."})
}

func getUsers(ctx *gin.Context) {
	usersList := database.GetUsers()
	ctx.JSON(http.StatusOK, usersList)
}
