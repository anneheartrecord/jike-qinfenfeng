package mysql

import (
	"golangstudy/jike/project1/models"

	"go.uber.org/zap"
)

func CheckRoomExist(roomname string) bool {
	var r models.Rooms
	DB.Where("room_name=?", roomname).First(&r)
	if r.ID != 0 {
		zap.L().Info("房间号已经存在")
		return true
	}
	return false
}
func CreateRoom(r models.Rooms) bool {
	DB.Create(&r)
	if r.ID == 0 {
		zap.L().Info("创建房间失败")
		return false
	}
	return true
}
func EnterRoom(roomname string) bool {
	var u models.Rooms
	DB.Where("room_name=?", roomname).First(&u)
	if u.ID == 0 {
		zap.L().Info("进入房间失败")
		return false
	}
	return true
}
func ChangeRoom(roomname, roomhost string) bool {
	var u models.Rooms
	DB.Model(&u).Where("room_name=?", roomname).Update("room_host", roomhost)
	if u.RoomHost != roomhost {
		zap.L().Info("update room host failed")
		return false
	}
	return true
}
func ExitRoom(roomname string) bool {
	var u models.Rooms
	DB.Delete(&u).Where("room_name=?", roomname)
	if u.ID != 0 {
		return false
	}
	return true
}
