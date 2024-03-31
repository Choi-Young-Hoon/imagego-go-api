package database

import (
	"errors"
	"gorm.io/gorm"
	"imagego-go-api/util"
)

func NewUser() User {
	return User{
		dbConnection: GeteDefaultDBConnection(),
	}
}

type User struct {
	gorm.Model
	UserId string `gorm:"unique;not null"`
	UserPw string `gorm:"not null"`

	dbConnection *DBConnection `gorm:"-"`
}

func (user *User) TableName() string {
	return "user"
}

func (user *User) Authorize(userId, password string) (bool, error) {
	err := user.FindByUserId(userId)
	if err != nil {
		return false, errors.New("사용자 계정이 존재하지 않습니다.")
	}

	if user.UserId != userId || user.UserPw != util.Sha512(password) {
		return false, errors.New("사용자 계정 또는 비밀번호가 일치하지 않습니다.")
	}
	return true, nil
}

func (user *User) FindById(id string) error {
	result := user.dbConnection.db.Where("id = ?", id).First(user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (user *User) FindByUserId(userId string) error {
	result := user.dbConnection.db.Where("user_id = ?", userId).First(user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (user *User) Create() error {
	user.UserPw = util.Sha512(user.UserPw)
	result := user.dbConnection.db.Create(user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
