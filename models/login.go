package models

type Login struct {
	UserName string `form:"username" json:"userName" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}
