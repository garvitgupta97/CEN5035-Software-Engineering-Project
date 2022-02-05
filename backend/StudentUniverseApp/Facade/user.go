package server

import (
	store "StudentUniverse/StudentUniverseApp/Facade/DTO"

	database "StudentUniverse/StudentUniverseApp/Facade/Database"

	"net/http"

	"github.com/gin-gonic/gin"
)

func signUp(ctx *gin.Context) {
	user := new(store.User)
	for _, u := range store.Users {
		if u.Email == user.Email {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"err": "User already exists."})
		}
	}
	if err := ctx.Bind(user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	database.InsertStudent(user.Email, user.Password)
	store.Users = append(store.Users, user)

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Signed up successfully.",
		"jwt": "123456789",
	})
}

func signIn(ctx *gin.Context) {
	user := new(store.User)
	if err := ctx.Bind(user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
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
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"err": "Sign in failed."})
}

func getUsers(ctx *gin.Context) {
	usersList := database.GetUsers()
	ctx.JSON(http.StatusOK, usersList)
}
