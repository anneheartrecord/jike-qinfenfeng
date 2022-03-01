package service

type UserService struct {
	Gender     string `json:"gender" form:"gender" binding:"required"`
	UserName   string `json:"user_name" form:"user_name" binding:"required"`
	Password   string `json:"password"   form:"password" binding:"required"`
	RePassword string `json:"re_password" form:"re_password" binding:"required,eqfield=Password"`
	Email      string `json:"email" form:"email" binding:"required"`
}
type NewUserService struct {
	Gender   string `json:"gender" form:"gender" binding:"required"`
	UserName string `json:"user_name" form:"user_name" binding:"required"`
	Password string `json:"password"   form:"password" binding:"required"`
	Code     string `json:"code" form:"code" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required"`
}
