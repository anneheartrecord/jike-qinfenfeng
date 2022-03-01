package main

import (
	"golangstudy/jike/project1/dao/mysql"
	"golangstudy/jike/project1/dao/redis"
	"golangstudy/jike/project1/logger"
	"golangstudy/jike/project1/pkg"
	"golangstudy/jike/project1/routers"
	"golangstudy/jike/project1/setting"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
)

func main() {
	// 加载配置(配置文件加载 远程加载)
	if err := setting.Init(); err != nil {

		zap.L().Error("setting init failed", zap.Error(err))
		return
	}
	// 初始化日志  大型项目必须使用日志
	if err := logger.Init(setting.Conf.LogConfig); err != nil {
		zap.L().Error("logger init failed", zap.Error(err))
	}
	defer zap.L().Sync() //缓冲区日志 追加到日志文件
	zap.L().Debug("logger init success")
	//初始化MySQL连接
	if err := mysql.Init(setting.Conf.MySQLConfig); err != nil {
		zap.L().Error("mysql init failed", zap.Error(err))
	}

	// 初始化Redis连接
	if err := redis.Init(setting.Conf.RedisConfig); err != nil {
		zap.L().Error("redis init failed", zap.Error(err))
	}
	defer redis.Close()
	if err := pkg.Init(setting.Conf.StartTime, setting.Conf.MachineID); err != nil {
		zap.L().Error("init snowflake failed", zap.Error(err))
		return
	}
	//初始化框架编译器
	if err := pkg.InitTrans("zh"); err != nil {
		zap.L().Error("init validator trans failed", zap.Error(err))
		return
	}
	r := routers.Init(setting.Conf)
	r.Run(":" + strconv.Itoa(setting.Conf.Port))
}
