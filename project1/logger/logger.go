package logger

import (
	"fmt"
	"golangstudy/jike/project1/setting"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var lg *zap.Logger //日志库 这里主要用的是zap

func Init(cfg *setting.LogConfig) (err error) { //初始化日志配置
	writerSyncer := getLogWriter(cfg.FileName, cfg.MaxSize, cfg.MaxBackups, cfg.MaxAge) //构造log writer
	encoder := getEncoder()                                                             //构造encoder
	var l = new(zapcore.Level)
	err = l.UnmarshalText([]byte(cfg.Level)) //日志记录level 这里是debug级别
	if err != nil {
		fmt.Println("unmarshal test failed", err)
		return
	}
	core := zapcore.NewCore(encoder, writerSyncer, l)
	lg = zap.New(core, zap.AddCaller())
	zap.ReplaceGlobals(lg) //替换zap包中lg的全局变量 调用zap.lg转换为zap.L()
	return
}
func getEncoder() zapcore.Encoder { //实现自定义的编码器
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder //配置一些属性
	return zapcore.NewJSONEncoder(encoderConfig)            //返回一个新的json数据格式编码器
}
func getLogWriter(filename string, maxSize, maxBackups, maxAge int) zapcore.WriteSyncer { //自定义的log writer 这里要用到lumberjack包来对日志进行切割
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackups,
		MaxAge:     maxAge,
	}
	return zapcore.AddSync(lumberJackLogger) //把缓存中的日志也进行输出
}
func GinLogger() gin.HandlerFunc { //这个函数接收gin框架默认的日志
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()
		cost := time.Since(start)
		zap.L().Info(path,
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("cost", cost),
		)
	}
}
func GinRecovery(stack bool) gin.HandlerFunc { //recover掉项目可能出现的panic 记录在日志文件里
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				var brokenPipe bool //检查是否为断掉的连接
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}
				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe { //断掉的进程就return
					zap.L().Error(c.Request.URL.Path, zap.Any("error", err), zap.String("request", string(httpRequest)))
					c.Error(err.(error))
					c.Abort()
					return
				}
				if stack {
					zap.L().Error("[Recovery from panic]")
					zap.Any("error", err)
					zap.String("request", string(httpRequest))
					zap.String("stack", string(debug.Stack()))
				} else {
					zap.L().Error("[Recovery from panic]")
					zap.Any("error", err)
					zap.String("request", string(httpRequest))
				}
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}

}
