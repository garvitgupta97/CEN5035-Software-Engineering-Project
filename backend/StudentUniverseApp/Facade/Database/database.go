package database

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type Users struct {
	Id       int    `gorm:"AUTO_INCREMENT" form:"id" json:"id"`
	Email    string `gorm:"not null" form:"email" json:"email"`
	Password string `gorm:"not null" form:"password" json:"password"`
}

type Profiles struct {
	ProfileId      int       `gorm:"column:id;primary_key;AUTO_INCREMENT"  form:"id" json:"id"`
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

func InsertUser(email string, password string) {
	db := InitializeDatabase()
	defer db.Close()

	user := new(Users)
	user.Email = email
	user.Password = password

	db.Create(&user)
}

func GetUsers() []string {
	var users []Users

	var userList []string

	db := InitializeDatabase()

	db.Select("Email, Password").Find(&users)

	for _, v := range users {
		userList = append(userList, v.Email, v.Password)
	}

	return userList
}

func InsertProfileData(profileId int, email string, name string, university string, profilePicture string, gender uint, birthDate time.Time, city string, state string, country string, bio string, createdAt time.Time, updatedAt time.Time) {

	db := InitializeDatabase()
	defer db.Close()
	profile := new(Profiles)
	profile.ProfileId = profileId
	profile.Name = name
	profile.Email = email
	profile.University = university
	profile.ProfilePicture = profilePicture
	profile.Gender = gender
	profile.BirthDate = birthDate
	profile.City = city
	profile.State = state
	profile.Country = country
	profile.Bio = bio
	profile.CreatedAt = createdAt
	profile.UpdatedAt = updatedAt

	// db.Clauses(clause.OnConflict{
	// 	Columns:   []clause.Column{{Name: "Email"}},                                                                                                                                                     // key colume
	// 	DoUpdates: clause.AssignmentColumns([]string{"ProfileId", "Email", "Name", "University", "ProfilePicture", "Gender", "BirthDate", "City", "State", "Country", "Bio", "CreatedAt", "UpdatedAt"}), // column needed to be updated
	// }).Create(&profile)

	if db.Model(&profile).Where("Email = ?", email).Updates(&profile).RowsAffected == 0 {
		db.Create(&profile)
	}
}

func GetProfiles() []Profiles {
	var profiles []Profiles
	db := InitializeDatabase()
	db.Find(&profiles)
	return profiles
}
