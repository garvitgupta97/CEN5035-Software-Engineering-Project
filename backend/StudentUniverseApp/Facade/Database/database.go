package database

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type Users struct {
	Id       int    `gorm:"AUTO_INCREMENT" form:"id" json:"id"`
	Email    string `gorm:"not null" form:"email" json:"email"`
	Password string `gorm:"not null" form:"password" json:"password"`
}

type Profiles struct {
	ProfileId      int       `gorm:"column:id; primary_key; AUTO_INCREMENT" form:"id" json:"id"`
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
	CreatedAt      time.Time `gorm:"not null column:created_at;default:CURRENT_TIMESTAMP" form:"created_at" json:"created_at"`
	UpdatedAt      time.Time `gorm:"not null column:updated_at;default:CURRENT_TIMESTAMP" form:"updated_at" json:"updated_at"`
}

type Thread struct {
	ThreadId    uuid.UUID `gorm:"column:thread_id; primary_key; AUTO_INCREMENT"`
	Title       string    `column:"title"`
	Description string    `column:"description"`
}

type Post struct {
	PostId        uuid.UUID `gorm:"column:post_id; primary_key; AUTO_INCREMENT"`
	UserId        int       `gorm:"column:user_id"`
	ThreadId      uuid.UUID `gorm:"column:thread_id"`
	Title         string    `gorm:"column:title"`
	Content       string    `gorm:"column:content"`
	Votes         int       `gorm:"column:votes"`
	CommentsCount int       `gorm:"column:comments_count"`
	ThreadTitle   string    `gorm:"column:thread_title"`
}

type Comment struct {
	CommentId uuid.UUID `gorm:"column:comment_id; primary_key; AUTO_INCREMENT"`
	PostId    uuid.UUID `gorm:"column:post_id"`
	Content   string    `gorm:"column:content"`
	Votes     int       `gorm:"column:votes"`
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
	if !db.HasTable(&Profiles{}) {
		db.CreateTable(&Profiles{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Profiles{})
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
	// profileTemp := new(Profiles)
	// profileTemp.ProfileId = profileId
	// profileTemp.Name = name
	// profileTemp.Email = email
	// profileTemp.University = university
	// profileTemp.ProfilePicture = profilePicture
	// profileTemp.Gender = gender
	// profileTemp.BirthDate = birthDate
	// profileTemp.City = city
	// profileTemp.State = state
	// profileTemp.Country = country
	// profileTemp.Bio = bio
	// profileTemp.CreatedAt = createdAt
	// profileTemp.UpdatedAt = updatedAt

	// db.Clauses(clause.OnConflict{
	// 	Columns:   []clause.Column{{Name: "Email"}},                                                                                                                                                     // key colume
	// 	DoUpdates: clause.AssignmentColumns([]string{"ProfileId", "Email", "Name", "University", "ProfilePicture", "Gender", "BirthDate", "City", "State", "Country", "Bio", "CreatedAt", "UpdatedAt"}), // column needed to be updated
	// }).Create(&profile)

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

func GetPostById(postId uuid.UUID) Post {
	post := new(Post)
	db := InitializeDatabase()
	defer db.Close()
	db.Model(&post).Where("PostId = ?", postId).First(&post)
	return *post
}


// func CreatePost(post Post) bool {
// 	db := InitializeDatabase()
// 	defer db.Close()
// 	if !isExistingUser(email) {
// 		user := new(Users)
// 		user.Email = email
// 		user.Password = password
// 		db.Create(&user)
// 		return true
// 	}
// 	return false
// }