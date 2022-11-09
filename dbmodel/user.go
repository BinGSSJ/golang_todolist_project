package dbmodel

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct { // user table
	gorm.Model
	UserName       string `gorm:unique`
	PasswordDigest string
}

// SetPwd 加密
func (user *User) SetPwd(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}
	user.PasswordDigest = string(bytes)
	return nil
}

// CheckPwd 验证密码
func (user *User) CheckPwd(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(password))
	return err == nil
}
