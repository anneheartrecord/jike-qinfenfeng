package controller

import (
	"golangstudy/jike/project1/logic"
	"golangstudy/jike/project1/models"
	"golangstudy/jike/project1/pkg"
	"golangstudy/jike/project1/service"

	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Register(c *gin.Context) {
	var p service.UserService
	if err := c.ShouldBind(&p); err != nil {
		zap.L().Error("register invalid ", zap.Error(err))
		err, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(200, models.Response{
				Status: 200,
				Error:  "参数格式错误",
			})
			return
		} else {
			c.JSON(200, models.Response{
				Error: err.Translate(pkg.Trans),
			})
			return
		}
	}
	if !logic.Register(p) {
		c.JSON(200, models.Response{
			Status: 200,
			Msg:    "用户名已被注册或邮箱已被绑定",
		})
		return
	}
	c.JSON(200, models.Response{
		Status: 200,
		Msg:    "注册成功",
	})
}
func Login(c *gin.Context) {
	var p service.UserService
	if err := c.ShouldBindJSON(&p); err != nil {
		zap.L().Error("login invalid ", zap.Error(err))
		c.JSON(200, models.Response{
			Status: 200,
			Msg:    "请求参数错误",
		})
		return

	}
	token, ok := logic.Login(p)
	if !ok {
		c.JSON(200, models.Response{
			Status: 200,
			Msg:    "用户名或密码错误",
		})
		return
	}
	c.JSON(200, models.Response{
		Status: 200,
		Data:   token,
		Msg:    "登录成功",
	})
}
func Update(c *gin.Context) {
	var p service.UserService
	if err := c.ShouldBindJSON(&p); err != nil {
		zap.L().Error("login invalid ", zap.Error(err))
		c.JSON(200, models.Response{
			Status: 200,
			Msg:    "请求参数错误",
		})
		return
	}
	_, ok := c.Get("username")
	if !ok {
		c.JSON(200, models.Response{
			Status: 200,
			Msg:    "您还未登录",
		})
		return
	}

	ok = logic.Update(p)
	if !ok {
		zap.L().Info("failed update")
		c.JSON(200, models.Response{
			Status: 200,
			Msg:    "修改用户名失败",
		})
		return
	}
	c.JSON(200, models.Response{
		Status: 200,
		Msg:    "修改用户名成功",
		Data:   p.UserName,
	})
}
func Exit(c *gin.Context) {
	var p service.UserService
	if err := c.ShouldBindJSON(&p); err != nil {
		zap.L().Error("login invalid ", zap.Error(err))
		c.JSON(200, models.Response{
			Status: 200,
			Msg:    "请求参数错误",
		})
		return
	}
	_, ok := c.Get("username")
	if !ok {
		c.JSON(200, models.Response{
			Status: 200,
			Msg:    "您还未登录",
		})
		return
	}
	ok = logic.Exit(p)
	if !ok {
		c.JSON(200, models.Response{
			Status: 200,
			Msg:    "退出系统失败",
		})
		return
	}
	c.JSON(200, models.Response{
		Status: 200,
		Msg:    "退出系统成功",
	})
}
func Changepwd(c *gin.Context) {
	var p service.UserService
	if err := c.ShouldBindJSON(&p); err != nil {
		zap.L().Error("login invalid ", zap.Error(err))
		c.JSON(200, models.Response{
			Status: 200,
			Msg:    "请求参数错误",
		})
		return
	}
	_, ok := c.Get("username")
	if !ok {
		c.JSON(200, models.Response{
			Status: 200,
			Msg:    "您还未登录",
		})
		return
	}
	ok = logic.Changepwd(p)
	if !ok {
		c.JSON(200, models.Response{
			Status: 200,
			Msg:    "修改密码失败",
		})
		return
	}
	c.JSON(200, models.Response{
		Status: 200,
		Msg:    "修改密码成功",
	})
}
func Forgetpwd(c *gin.Context) {
	var p service.UserService
	if err := c.ShouldBindJSON(&p); err != nil {
		zap.L().Error("login invalid ", zap.Error(err))
		c.JSON(200, models.Response{
			Status: 200,
			Msg:    "请求参数错误",
		})
		return
	}
	_, ok := c.Get("username")
	if !ok {
		c.JSON(200, models.Response{
			Status: 200,
			Msg:    "您还未登录",
		})
		return
	}
	code, ok := logic.Forgetpwd(p)
	if !ok {
		c.JSON(200, models.Response{
			Status: 200,
			Msg:    "没有该邮箱",
		})
		return
	}
	c.JSON(200, models.Response{
		Status: 200,
		Data:   code,
		Msg:    "验证码发放成功",
	})
}
func Verify(c *gin.Context) {
	var p service.NewUserService
	if err := c.ShouldBindJSON(&p); err != nil {
		zap.L().Error("login invalid ", zap.Error(err))
		c.JSON(200, models.Response{
			Status: 200,
			Msg:    "请求参数错误",
		})
		return
	}
	_, ok := c.Get("username")
	if !ok {
		c.JSON(200, models.Response{
			Status: 200,
			Msg:    "您还未登录",
		})
		return
	}
	ok = logic.Verify(p)
	if !ok {
		c.JSON(200, models.Response{
			Status: 200,
			Msg:    "验证码错误",
		})
		return
	}
	c.JSON(200, models.Response{
		Status: 200,
		Msg:    "改密成功",
	})
}
