package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type Users struct {
	Id       int    `gorm:"AUTO_INCREMENT" form:"id" json:"id"`
	Email    string `gorm:"not null" form:"email" json:"email"`
	Password string `gorm:"not null" form:"password" json:"password"`
}

func InitializeDatabase() *gorm.DB {
	// Openning file
	db, err := gorm.Open("sqlite3", "./data.db")
	// Display SQL queries
	db.LogMode(true)

	// Error
	if err != nil {
		panic(err)
	}
	// Creating the table
	if !db.HasTable(&Users{}) {
		db.CreateTable(&Users{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Users{})
	}

	return db
}

func InsertUser(email string, password string) {
	db := InitializeDatabase()
	defer db.Close()

	user := new(Users)
	user.Email = email
	user.Password = password

	db.Create(&user)
}

// func GetUsers() []string {
// 	var users []Users

// 	var userList []string

// 	db := InitializeDatabase()

// 	db.Select("Email").Find(&users)

// 	for _, v := range users {
// 		userList = append(userList, v.Email)
// 	}

// 	return userList
// }

func GetUsers() []Users {
	var userList []Users

	//var userList []string

	db := InitializeDatabase()
	//db.Raw("SELECT id, Email, Password FROM users").Scan(&users)

	db.Table("Users").Select("Id, Email").Find(&userList)

	return userList
}
