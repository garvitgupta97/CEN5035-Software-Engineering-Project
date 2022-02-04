package gin_router
import (
    api "backend/gin_api"

    "github.com/gin-gonic/gin"
)
func InitRouter() {
    gin.SetMode(gin.ReleaseMode)
    //Create a routing handler using the Default method of gin
    router := gin.Default()
    //Set default route return when visiting a wrong website
    router.NoRoute(api.NotFound)
    //Use the Group function provided by the following gin to Group different API s
    v1 := router.Group("admin")
    {
        v1.GET("/register", api.Register)
        v1.GET("/login", api.Login)
    }
    //Listening server port
    router.Run(":8080")
}