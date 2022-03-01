package controller

import (
	"golangstudy/jike/project1/logic"
	"golangstudy/jike/project1/models"
	"golangstudy/jike/project1/pkg"
	"golangstudy/jike/project1/service"

	"github.com/gorilla/websocket"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

func CreateRoom(c *gin.Context) {
	var p service.RoomService
	if err := c.ShouldBindJSON(&p); err != nil {
		zap.L().Error("create invalid ", zap.Error(err))
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
	v, ok := c.Get("username")
	p.RoomHost = v.(string)
	if !ok {
		c.JSON(200, models.Response{
			Status: 200,
			Msg:    "您还未登录",
		})
		return
	}

	ok = logic.CreateRoom(p)
	if !ok {
		c.JSON(200, models.Response{
			Status: 200,
			Msg:    "已经有该房间",
		})
		return
	}
	conn, err := (&websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}).Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		zap.L().Error("ws upgrade failed", zap.Error(err))
	}
	client := &models.Client{ //用户实例
		ID:     p.RoomHost,
		SendID: p.RoomName,
		Socket: conn,
		Send:   make(chan []byte),
	}
	models.Manager.Register <- client
	go client.Read()
	//go client.Write()
	c.JSON(200, models.Response{
		Status: 200,
		Msg:    "创建房间成功",
	})
}
func EnterRoom(c *gin.Context) {
	var p service.RoomService
	if err := c.ShouldBindJSON(&p); err != nil {
		zap.L().Error("create invalid ", zap.Error(err))
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
	ok := logic.EnterRoom(p)
	if !ok {
		c.JSON(200, models.Response{
			Status: 200,
			Msg:    "进入房间失败",
		})
	}
	c.JSON(200, models.Response{
		Status: 200,
		Msg:    "进入房间成功",
	})
}
func ChangeRoom(c *gin.Context) {
	var p service.RoomService
	if err := c.ShouldBindJSON(&p); err != nil {
		zap.L().Error("create invalid ", zap.Error(err))
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
	ok := logic.ChangeRoom(p)
	if !ok {
		c.JSON(200, models.Response{
			Status: 200,
			Msg:    "修改房间失败",
		})
	}
	c.JSON(200, models.Response{
		Status: 200,
		Msg:    "修改房间成功",
	})
}
func ExitRoom(c *gin.Context) {
	var p service.RoomService
	if err := c.ShouldBindJSON(&p); err != nil {
		zap.L().Error("create invalid ", zap.Error(err))
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
	ok := logic.ExitRoom(p)
	if !ok {
		c.JSON(200, models.Response{
			Status: 200,
			Msg:    "退出房间失败",
		})
	}
	c.JSON(200, models.Response{
		Status: 200,
		Msg:    "退出房间成功",
	})
}
