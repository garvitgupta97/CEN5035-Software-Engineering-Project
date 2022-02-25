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
	// for _, u := range store.Users {
	// 	if u.Email == user.Email {
	// 		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"err": "User already exists."})
	// 	}
	// }
	if err := ctx.Bind(user); err != nil {

		var verr validator.ValidationErrors
		if errors.As(err, &verr) {
			ctx.JSON(http.StatusBadRequest, gin.H{"errors": SimpleErrorMsg(verr)})
			return
		}

		log.Info().Err(err).Msg("unable to bind")

		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Signup Fail": err.Error()})
		return
	}
	// encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	// if err != nil {
	// 	ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Signup Fail": err.Error()})
	// 	return
	// }

	if !database.InsertUser(user.Email, user.Password) {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Signup Fail": "User Already Exists"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Signed up successfully.",
		"jwt": "123456789",
	})
}

func SimpleErrorMsg(verr validator.ValidationErrors) map[string]string {
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
	user := new(database.Users)
	if err := ctx.Bind(user); err != nil {
		var verr validator.ValidationErrors
		if errors.As(err, &verr) {
			ctx.JSON(http.StatusBadRequest, gin.H{"errors": SimpleErrorMsg(verr)})
			return
		}
		log.Info().Err(err).Msg("unable to bind")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Sign in failed": err.Error()})
		return
	}
	//encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	// if err != nil {
	// 	ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Sign in failure": err.Error()})
	// }
	isVerifiedUser := database.IsVerifiedUser(user.Email, user.Password) //string(encryptedPassword))

	if isVerifiedUser {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "Signed in successfully.",
			"jwt": "123456789",
		})
		return
	}
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Sign in failed": "User not found"})
}

func getUsers(ctx *gin.Context) {
	usersList := database.GetUsers()
	ctx.JSON(http.StatusOK, usersList)
}

func getProfiles(ctx *gin.Context) {
	usersList := database.GetProfiles()
	ctx.JSON(http.StatusOK, usersList)
}

func updateProfile(ctx *gin.Context) {
	profile := new(database.Profiles)
	if err := ctx.Bind(profile); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Errors": "Input error"})
	}
	isProfileUpdated := database.UpsertProfile(profile)
	fmt.Println("Final Ans: ", isProfileUpdated)
	if isProfileUpdated {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "Profile updated successfully.",
		})
		return
	}
	ctx.JSON(http.StatusBadRequest, gin.H{
		"Error": "Profile updatation failed",
	})

}

// func signUp2(ctx *gin.Context) {
// 	user := new(store.User)
// 	err := ctx.Bind(user)
// 	// if err := ctx.Bind(user); err != nil {

// 	// 	var verr validator.ValidationErrors
// 	// 	if errors.As(err, &verr) {
// 	// 		ctx.JSON(http.StatusBadRequest, gin.H{"errors": SimpleErrorMsg(verr)})
// 	// 		return
// 	// 	}

// 	// 	log.Info().Err(err).Msg("unable to bind")

// 	// 	ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Signup Fail": err.Error()})
// 	// 	return
// 	// }

// 	fmt.Println(err)
// 	fmt.Println("User " + user.Email + "Pass: " + user.Password)
// 	ctx.JSON(http.StatusOK, gin.H{
// 		"msg": "Signed up successfully.",
// 	})
// }
