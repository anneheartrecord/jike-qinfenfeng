package logic

import (
	"golangstudy/jike/project1/dao/mysql"
	"golangstudy/jike/project1/dao/redis"
	"golangstudy/jike/project1/models"
	"golangstudy/jike/project1/pkg"
	"golangstudy/jike/project1/service"

	"go.uber.org/zap"
)

func Register(u service.UserService) bool {
	exist := mysql.CheckUserExist(u.UserName)
	if exist {
		zap.L().Info("user exist")
		return false
	}
	exist = mysql.CheckEmailExist(u.Email)
	if exist {
		zap.L().Info("email binded")
		return false
	}
	userID := pkg.GenID()

	U := models.Users{
		UID:        userID,
		UserName:   u.UserName,
		Password:   u.Password,
		RePassword: u.RePassword,
		Email:      u.Email,
		Gender:     u.Gender,
	}
	mysql.InsertUser(U)
	return true
}
func Login(u service.UserService) (string, bool) {
	exist := mysql.CheckUserExist(u.UserName)
	if !exist {
		zap.L().Info("user not exist")
		return "", false
	}
	equal := mysql.CheckPassword(u.UserName, u.Password)
	if !equal {
		zap.L().Info("wrong password")
		return "", false
	}
	token, err := pkg.GenToken(u.UserName)
	if err != nil {
		zap.L().Error("gen token failed", zap.Error(err))
		return "", false
	}
	return token, true
}
func Update(u service.UserService) bool {
	ok := mysql.CheckUserExist(u.UserName)
	if ok {
		zap.L().Info("the username exist")
		return false
	}
	ok = mysql.UpdateUsername(u.Email, u.UserName)
	if !ok {
		zap.L().Info("update username failed")
		return false
	}
	return true
}
func Exit(u service.UserService) bool {
	ok := mysql.DeleteUser(u.UserName)
	if !ok {
		zap.L().Info("delete user failed")
		return false
	}
	return true
}
func Changepwd(u service.UserService) bool {
	ok := mysql.Changepwd(u.Email, u.Password)
	if !ok {
		zap.L().Info("change password failed")
		return false
	}
	return true
}
func Forgetpwd(u service.UserService) (string, bool) {
	ok := mysql.CheckEmailExist(u.Email)
	if !ok {
		zap.L().Info("not exist the email")
		return "", false
	}
	code := redis.AuthCode()
	return code, true
}
func Verify(u service.NewUserService) bool {
	return redis.GetAuthCode(u.Code)
}
