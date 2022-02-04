package data
//Structure used to store user information, Id,Name,Passwd
type User struct {
    Id     int
    Name   string
    Passwd string
}
//Slices for storing users
var Slice []User
//Map for temporary storage of user login information
var State = make(map[string]interface{})
