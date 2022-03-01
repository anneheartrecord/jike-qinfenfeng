package redis

import (
	"golangstudy/jike/project1/pkg"
	"time"

	"go.uber.org/zap"
)

func AuthCode() string {
	code := pkg.RandomCode()
	rdb.Set("code", code, 5*time.Minute)
	return code
}
func GetAuthCode(code string) bool {
	authcode, err := rdb.Get("code").Result()
	if err != nil {
		zap.L().Error("get code failed", zap.Error(err))
		return false
	}
	if code != authcode {
		zap.L().Info("authcode not true")
		return false
	}
	return true
}
