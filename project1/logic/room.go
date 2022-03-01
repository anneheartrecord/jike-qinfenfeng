package logic

import (
	"golangstudy/jike/project1/dao/mysql"
	"golangstudy/jike/project1/models"
	"golangstudy/jike/project1/pkg"
	"golangstudy/jike/project1/service"

	"go.uber.org/zap"
)

func CreateRoom(u service.RoomService) bool {
	exist := mysql.CheckRoomExist(u.RoomName)
	if exist {
		zap.L().Info("room  exist")
		return false
	}
	roomID := pkg.GenID()
	U := models.Rooms{
		RoomID:   roomID,
		RoomName: u.RoomName,
		RoomCap:  u.RoomCap,
		RoomHost: u.RoomHost,
	}
	mysql.CreateRoom(U)
	return true
}
func EnterRoom(u service.RoomService) bool {
	exist := mysql.CheckRoomExist(u.RoomName)
	if !exist {
		zap.L().Info("room  not exist")
		return false
	}
	mysql.EnterRoom(u.RoomName)
	return true
}
func ChangeRoom(u service.RoomService) bool {
	exist := mysql.CheckRoomExist(u.RoomName)
	if !exist {
		zap.L().Info("room  not exist")
		return false
	}
	mysql.ChangeRoom(u.RoomName, u.RoomHost)
	return true
}
func ExitRoom(u service.RoomService) bool {
	exist := mysql.CheckRoomExist(u.RoomName)
	if !exist {
		zap.L().Info("room  not exist")
		return false
	}
	mysql.ExitRoom(u.RoomName)
	return true
}
