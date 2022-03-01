package mysql

import (
	"fmt"
	"golangstudy/jike/project1/models"
	"golangstudy/jike/project1/setting"

	"go.uber.org/zap"

	_ "github.com/go-sql-driver/mysql"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func Init(cfg *setting.MySQLConfig) (err error) { //初始化MySQL连接
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true",
		cfg.Root,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DbName,
	)
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		zap.L().Error("connect mysql failed", zap.Error(err))
	}
	MysqlDB := DB.DB()
	MysqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	MysqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	DB.AutoMigrate(&models.Users{})
	DB.AutoMigrate(&models.Rooms{})
	return
}
