package model

import (
	"github.com/jinzhu/gorm"
	"github.com/ohmyray/my-blog/utils"
)

type User struct {
	gorm.Model
	Username string `gorm:username "type:varchar(20);not null " json:"username"`
	Password string `gorm:password "type:varchar(20);not null " json:"password"`
}

func CheckUser(username string) int {
	var user User
	connection.Select("id").Where("username = ?", username).First(&user)
	if user.ID > 0 {
		// 用户已被使用
		return utils.USER_EXIST_USED
	}
	// 用户已存在
	return utils.USER_OK
}

func AddUser(user *User) int {
	err = connection.Create(&user).Error
	if err != nil {
		return utils.USER_REGISTER_FAIL
	}
	return utils.USER_REGISTER_SUCCESS
}

func UserInfo(id string) (data *User, code int) {
	var user User
	connection.Where("id = ?", id).First(&user)
	if user.ID == 0 {
		code = utils.USER_NO_EXIST
	}
	code = utils.USER_EXIST
	return &user, code
}
