package store

type User struct {
	Email    string `binding:"required,min=5,max=30, regexp=^[_A-Za-z0-9+-]+(\\.[_A-Za-z0-9-]+)*@[A-Za-z0-9-]+(\\.[A-Za-z0-9]+)*(\\.[A-Za-z]{2\\,})$"`
	Password string `binding:"required,min=7,max=32"`
}

var Users []*User
