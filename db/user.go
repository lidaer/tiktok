package db

import (
	"fmt"
	"sync/atomic"
)

// User : 用户表model
type User struct {
	ID         int64
	UserInfoId int64
	username   string
	password   string
}

func (u User) TableName() string {
	return "user_logins"
}

var userIdSequence = int64(1)

// AddUser addUser: 添加用户
func AddUser(username string, passwd string) int64 {
	atomic.AddInt64(&userIdSequence, 1)
	user := User{username: username, password: passwd, UserInfoId: userIdSequence}
	err := DB.Create(&user).Error
	if err != nil {
		fmt.Println("Failed to insert, err:" + err.Error())
		return -100
	}
	return userIdSequence
}

// UserExist 查询用户名是否重复
func UserExist(username string) bool {
	user := User{username: username}
	result := DB.First(user)
	if result.Error != nil {
		fmt.Println("Failed to insert, err:" + result.Error.Error())
		return false
	}
	if result.RowsAffected == 0 {
		return false
	}
	return true
}
