package model

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/snowitty/fabler/internal/database"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type User struct {
	GORMBase
	Mail string `json:"mail" gorm:"type:varchar(100);unique_index"`
	Password string `json:"password"`
	Authorized int `json:"authorized"`
	Type int `json:"type"`
	Profile Profile `json:"profile" gorm:"ForeignKey:ProfileID"`
	ProfileID int `json:"profileID"`
}

func (u *User) Get() (user User, err error){
	if err = database.DB.Where(u).Preload("Profile").First(&user).Error;err != nil{
		log.Print(err)
	}

	return
}

func (u *User) Create() (ra int64, err error){
	if err = database.DB.Create(&u).Error;err != nil{
		ra = -1
		log.Print(err)
	}else {
		ra = 1
	}
	return
}

func (u *User) Update() (ra int64, err error){
	if err = database.DB.Model(&u).Update(u).Error;err != nil{
		ra = -1
		log.Print(err)
	}else {
		ra = 1
	}
	return
}

func (u *User) Delete() (ra int64, err error){
	if err = database.DB.Delete(&u).Error;err != nil{
		ra = -1
		log.Print(err)
	}else {
		ra = 1
	}
	return
}

func (u *User) GetList(limit int,offset int) (user []User, err error){
	if err = database.DB.Offset(offset).Limit(limit).Find(&user).Error;err != nil{
		log.Print(err)
	}
	return
}

func (u *User) GetCounts() (counts int, err error){
	if err = database.DB.Model(&User{}).Count(&counts).Error;err != nil{
		log.Print(err)
	}

	return
}

func (u *User) BeforeDelete(tx *gorm.DB) (err error){
	if u.ID == 1{
		err = errors.New("Can not Remove Admin")
	}
	return
}

func (u *User) AfterCreate(tx *gorm.DB) (err error){
	bytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil{
		log.Print(err)
		return
	}
	if err = tx.Model(&u).UpdateColumn("password", string(bytes)).Error; err != nil{
		log.Print(err)
	}

	return
}

