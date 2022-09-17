package Models

type Login struct {
	Username string `form:"username" json:"userName" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}
