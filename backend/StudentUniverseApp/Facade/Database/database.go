package database

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type Users struct {
	Id       int    `gorm:"AUTO_INCREMENT" form:"id" json:"id"`
	Email    string `gorm:"not null" form:"email" json:"email" binding:"required,min=5,max=30,email"`
	Password string `gorm:"not null" form:"password" json:"password" binding:"required,min=8,max=30"`
}

type Profiles struct {
	ProfileId      int       `gorm:"AUTO_INCREMENT; PRIMARY_KEY; column:id;" form:"id" json:"id"`
	Email          string    `gorm:"not null" form:"email" json:"email"`
	Name           string    `gorm:"not null column:user_name;NOT NULL" form:"username" json:"username"`
	University     string    `gorm:"not null column:university" form:"university" json:"university"`
	ProfilePicture string    `gorm:"not null column:profile_picture" form:"profile_picture" json:"profile_picture"`
	Gender         uint      `gorm:"not null column:gender" form:"gender" json:"gender"`
	BirthDate      time.Time `gorm:"not null column:birth_date" form:"birth_date" json:"birth_date"`
	City           string    `gorm:"not null column:city" form:"city" json:"city"`
	State          string    `gorm:"not null column:state" form:"state" json:"state"`
	Country        string    `gorm:"not null column:country" form:"country" json:"country"`
	Bio            string    `gorm:"not null column:bio" form:"bio" json:"bio"`
	CreatedAt      time.Time `gorm:"not null column:created_at;"`
	UpdatedAt      time.Time `gorm:"not null column:updated_at;"`
}

type Thread struct {
	ThreadId    int    `gorm:"column:thread_id; primary_key; AUTO_INCREMENT"`
	Title       string `column:"title"`
	Description string `column:"description"`
}

type Post struct {
	PostId        int       `gorm:"column:post_id; primary_key; AUTO_INCREMENT"`
	UserId        int       `gorm:"column:user_id"`
	ThreadId      int       `gorm:"column:thread_id"`
	Title         string    `gorm:"column:title"`
	Content       string    `gorm:"column:content"`
	Votes         int       `gorm:"column:votes"`
	CommentsCount int       `gorm:"column:comments_count"`
	CreatedAt     time.Time `gorm:"not null column:created_at;"`
	UpdatedAt     time.Time `gorm:"not null column:updated_at;"`
	ThreadTitle   string    `gorm:"column:thread_title"`
}

type PostVotes struct {
	Id        int    `gorm:"column:id; primary_key; AUTO_INCREMENT"`
	PostId    int    `gorm:"column:post_id"`
	UserEmail string `gorm:"column:user_email"`
	VoteValue int    `gorm:"column:votes"`
}

type Comment struct {
	CommentId int       `gorm:"column:comment_id; primary_key; AUTO_INCREMENT"`
	UserId    int       `gorm:"column:user_id"`
	PostId    int       `gorm:"column:post_id"`
	Content   string    `gorm:"column:content"`
	CreatedAt time.Time `gorm:"not null column:created_at;"`
	UpdatedAt time.Time `gorm:"not null column:updated_at;"`
}

type CommentPost struct {
	Post    Post
	Comment []int
}

type AllPosts struct {
	PostId        int
	UserId        int
	ThreadId      int
	Title         string
	Content       string
	Votes         int
	CommentsCount int
	CreatedAt     time.Time
	UpdatedAt     time.Time
	ThreadTitle   string
	Email         string
}

type AllComments struct {
	CommentId int
	UserId    int
	PostId    int
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
	Email     string
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
	//db.AutoMigrate(&Post{})

	if !db.HasTable(&Users{}) {
		db.CreateTable(&Users{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Users{})
	}
	if !db.HasTable(&Profiles{}) {
		db.CreateTable(&Profiles{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Profiles{})
	}
	if !db.HasTable(&Post{}) {
		db.CreateTable(&Post{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Post{})
	}

	if !db.HasTable(&PostVotes{}) {
		db.CreateTable(&PostVotes{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&PostVotes{})
	}

	return db
}

func IsVerifiedUser(email string, password string) bool {
	db := InitializeDatabase()
	defer db.Close()
	user := new(Users)
	return db.Model(&user).Where("Email = ? AND Password = ?", email, password).First(&user).RowsAffected != 0
}

func InsertUser(email string, password string) bool {
	db := InitializeDatabase()
	defer db.Close()
	if !isExistingUser(email) {
		user := new(Users)
		user.Email = email
		user.Password = password
		db.Create(&user)
		return true
	}
	return false
}

func GetUsers() []Users {
	var users []Users
	db := InitializeDatabase()
	defer db.Close()
	db.Select("Id, Email").Find(&users)
	return users
}

func isExistingUser(email string) bool {
	user := new(Users)
	db := InitializeDatabase()
	defer db.Close()
	return db.Model(&user).Where("Email = ?", email).First(&user).RowsAffected != 0
}

//func InsertProfileData(profileId int, email string, name string, university string,
//profilePicture string, gender uint, birthDate time.Time, city string, state string, country string,
//bio string, createdAt time.Time, updatedAt time.Time) bool {
func UpsertProfile(profile *Profiles) bool {
	db := InitializeDatabase()
	defer db.Close()

	if db.Model(&profile).Where("Email = ?", profile.Email).Updates(&profile).RowsAffected == 0 {
		if isExistingUser(profile.Email) {
			return db.Create(&profile).Error == nil
		}
		return false
	}
	return true
}

func GetProfiles() []Profiles {
	var profiles []Profiles
	db := InitializeDatabase()
	defer db.Close()
	db.Find(&profiles)
	return profiles
}

func TestQuery(query string, email string, password string) bool {
	db := InitializeDatabase()
	defer db.Close()
	var user []Users
	ans := db.Raw(query, email, password).First(&user).RowsAffected != 0

	return ans
}

func GetPostById(postId int) Post {
	post := new(Post)
	db := InitializeDatabase()
	defer db.Close()
	db.Model(&post).Where("post_id = ?", postId).First(&post)
	fmt.Print(postId, *post)
	return *post
}

func CreatePost(post Post) bool {
	db := InitializeDatabase()
	defer db.Close()
	//fmt.Println(db.Create(&post).Error)
	return db.Create(&post).Error == nil
}

func CreateComment(comment Comment) bool {
	db := InitializeDatabase()
	defer db.Close()
	//fmt.Println(db.Create(&post).Error)
	return db.Create(&comment).Error == nil
}

func GetAllPosts() []AllPosts {
	//var allPosts Post

	//allPosts := AllPosts{}
	var allPosts []AllPosts
	db := InitializeDatabase()
	defer db.Close()
	//db.Model(&allPosts).Preload("Users").Find(&allPosts)
	db.Table("posts").Select("posts.*, users.email").Joins("inner join users on posts.user_id = users.id").Find(&allPosts)
	return allPosts
}

func GetCommentsByPosts(postId int) []AllComments {

	var allComments []AllComments
	db := InitializeDatabase()
	defer db.Close()
	//db.Model(&allPosts).Preload("Users").Find(&allPosts)
	db.Table("comments").Where("post_id = ?", postId).Select("comments.*, users.email").Joins("inner join users on comments.user_id = users.id").Find(&allComments)
	return allComments
}

func AddPostVote(postVotes PostVotes) bool {
	db := InitializeDatabase()
	defer db.Close()
	//fmt.Println(db.Create(&post).Error)
	return db.Create(&postVotes).Error == nil
}
