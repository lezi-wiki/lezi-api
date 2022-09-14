package model

import (
	"encoding/gob"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string  `json:"username" xml:"username" bson:"username" gorm:"not null;uniqueIndex"`
	Password string  `json:"-" xml:"-" bson:"-" gorm:"not null"`
	Email    string  `json:"email" xml:"email" bson:"email" gorm:"not null;uniqueIndex"`
	Avatar   *string `json:"avatar" xml:"avatar" bson:"avatar"`

	Texts []Text `json:"texts" xml:"texts" bson:"texts" gorm:"foreignKey:UserRefer"`
}

func init() {
	gob.Register(&User{})
}

type UserService interface {
	Get(user User) (*User, error)
	List(user User) ([]User, error)
	ListAll() ([]User, error)
	CreateUser(user User) (*User, error)
	UpdateUser(id uint, user User) (*User, error)
	CheckLogin(username, password string) (*User, bool)
	DeleteUser(id uint) error
	Count() int64
	Exists(user User) bool
}

type UserServiceImpl struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) UserService {
	return &UserServiceImpl{db: db}
}

func (u *UserServiceImpl) Get(user User) (*User, error) {
	err := u.db.Model(&User{}).Where(&user).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *UserServiceImpl) List(user User) ([]User, error) {
	var users []User
	err := u.db.Model(&User{}).Where(&user).Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (u *UserServiceImpl) ListAll() ([]User, error) {
	return u.List(User{})
}

func (u *UserServiceImpl) CreateUser(user User) (*User, error) {
	err := u.db.Model(&User{}).Create(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *UserServiceImpl) UpdateUser(id uint, user User) (*User, error) {
	err := u.db.Model(&User{}).Where("id = ?", id).Updates(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *UserServiceImpl) CheckLogin(username, password string) (*User, bool) {
	var user User
	err := u.db.Model(&User{}).Where("username = ? AND password = ?", username, password).First(&user).Error
	if err != nil {
		return nil, false
	}

	return &user, true
}

func (u *UserServiceImpl) DeleteUser(id uint) error {
	err := u.db.Model(&User{}).Where("id = ?", id).Delete(&User{}).Error
	if err != nil {
		return err
	}

	return nil
}

func (u *UserServiceImpl) Count() int64 {
	var count int64 = 0
	u.db.Model(&User{}).Count(&count)
	return count
}

func (u *UserServiceImpl) Exists(user User) bool {
	var count int64 = 0
	u.db.Model(&User{}).Where(&user).Count(&count)
	return count > 0
}
