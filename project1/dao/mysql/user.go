package mysql

import (
	"golangstudy/jike/project1/models"
	"golangstudy/jike/project1/pkg"

	"go.uber.org/zap"
)

func CheckUserExist(username string) bool {
	var u models.Users
	DB.Where("user_name=?", username).First(&u)
	if u.ID != 0 {
		zap.L().Info("用户名已经存在")
		return true
	}
	return false
}
func CheckEmailExist(email string) bool {
	var u models.Users
	DB.Where("email=?", email).First(&u)

	if u.ID != 0 {
		zap.L().Info("邮箱已经被绑定")
		return true
	}
	return false
}
func InsertUser(user models.Users) {
	user.Password = pkg.EncryptPassword(user.Password)
	user.RePassword = pkg.EncryptPassword(user.RePassword)
	DB.Create(&user)
}
func CheckPassword(username, password string) bool {
	var u models.Users
	DB.Where("user_name=?", username).First(&u)
	DB.Where("password=?", password).First(&u)
	if pkg.EncryptPassword(password) == u.Password {
		return true
	}
	return false
}
func UpdateUsername(email, username string) bool {
	var u models.Users
	DB.Model(&u).Where("email=?", email).Update("user_name", username)
	if u.UserName == username {
		zap.L().Info("update username success")
		return true
	}
	return false
}
func DeleteUser(username string) bool {
	var u models.Users
	DB.Delete(&u).Where("user_name=?", username)
	if u.ID != 0 {
		return false
	}
	return true
}
func Changepwd(email, password string) bool {
	var u models.Users
	password = pkg.EncryptPassword(password)
	DB.Model(&u).Where("email=?", email).Update("password", password)
	DB.Model(&u).Where("email=?", email).Update("re_password", password)
	return true
}
