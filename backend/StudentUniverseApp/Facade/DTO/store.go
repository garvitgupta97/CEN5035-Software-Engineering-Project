package store

type User struct {
	// Email    string `binding:"required,min=5,max=30, regexp=^[_A-Za-z0-9+-]+(\\.[_A-Za-z0-9-]+)*@[A-Za-z0-9-]+(\\.[A-Za-z0-9]+)*(\\.[A-Za-z]{2\\,})$"`
	Email    string `binding:"required,min=5,max=30,email"`
	Password string `binding:"required,min=7,max=32"`
}
type Post struct {
	Title   string `binding:"required,min=10,max=30"`
	Content string `binding:"required,min=10,max=100"`
	Email   string `binding:"required,min=5,max=30,email"`
}

var Users []*User
var Posts []*Post
