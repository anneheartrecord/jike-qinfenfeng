package setting

import (
	"fmt"

	"go.uber.org/zap"

	"github.com/fsnotify/fsnotify"

	"github.com/spf13/viper"
)

var Conf = new(ChatRoomConfig) //读取配置文件 这里用的是viper

type ChatRoomConfig struct {
	Name         string `mapstructure:"name"`
	Port         int    `mapstructure:"port"`
	StartTime    string `mapstructure:"start_time"`
	MachineID    int64  `mapstructure:"machine_id"`
	*LogConfig   `mapstructure:"log"`
	*MySQLConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
}
type LogConfig struct {
	Level      string `mapstructure:"level"`
	FileName   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}
type MySQLConfig struct {
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	Root         string `mapstructure:"root"`
	Password     string `mapstructure:"password"`
	DbName       string `mapstructure:"dbname"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}
type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	DB       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"pool_size"`
}

func Init() (err error) {
	//这里我们使用viper包来加载config文件配置
	viper.SetConfigName("config")
	viper.AddConfigPath(".")   //指定相对路径
	err = viper.ReadInConfig() //读取配置
	if err != nil {
		fmt.Println("viper read config failed", err)
		return
	}
	if err = viper.Unmarshal(Conf); err != nil {
		zap.L().Error("viper unmarshal failed", zap.Error(err))
		return
	}
	viper.WatchConfig()                            //viper包里的热加载
	viper.OnConfigChange(func(in fsnotify.Event) { //热加载之后再反序列化一次 这样就可以保存配置了
		if err := viper.Unmarshal(Conf); err != nil {
			zap.L().Error("viper unmarshal failed", zap.Error(err))
			return
		}
	})
	return
}
