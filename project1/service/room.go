package service

type RoomService struct {
	RoomName string `json:"room_name" form:"room_name"`
	RoomCap  int    `json:"room_cap" form:"room_cap"`
	RoomHost string `json:"room_host" form:"room_host"`
}
