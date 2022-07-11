package model

type Login struct {
	User     string `json:"userName" binding:"required"`
	Password string `json:"userPassword" binding:"required"`
}
