package gin_regAndlog
import (
    . "backend/data"
    Func "backend/func_judge"
    "net/http"

    "github.com/gin-gonic/gin"
)
//register
func Register(c *gin.Context) {
    //Get user name and password
    name := c.Request.FormValue("Name")
    passwd := c.Request.FormValue("Passwd")
    //Judge whether the user exists
    //Output state 1 present
    //No create user exists, save password and user name
    Bool := Func.IsExist(name)
	
    if Bool {
        //Registration status
        State["state"] = 1
        State["text"] = "This user already exists!"
    } else {
        //Add user if user does not exist
        AddStruct(name, passwd)
        State["state"] = 1
        State["text"] = "Registration succeeded!"
    }

    //Return the status code and registration status to the client
    c.String(http.StatusOK, "%v", State)
}
//Sign in
func Login(c *gin.Context) {
    name := c.Request.FormValue("Name")
    passwd := c.Request.FormValue("Passwd")
    //Judge whether the user exists first, and then judge whether the password is correct
    Bool := Func.IsExist(name)
    if Bool {
        Bool_Pwd := Func.IsRight(name, passwd)
        if Bool_Pwd {
            State["state"] = 1
            State["text"] = "Login succeeded!"
        } else {
            State["state"] = 0
            State["text"] = "Wrong password!"
        }
    } else {
        State["state"] = 2
        State["text"] = "Login failed! This user is not registered!"
    }

    c.String(http.StatusOK, "%v", State)
}
//Set default route return when visiting a wrong website
func NotFound(c *gin.Context) {
    c.JSON(http.StatusNotFound, gin.H{
        "status": 404,
        "error":  "404 ,page not exists!",
    })
}
//Add user
func AddStruct(name string, passwd string) {
    var user User
    user.Name = name
    user.Passwd = passwd
    user.Id = len(Slice) + 1
    Slice = append(Slice, user)
}