package main

import (
	"github.com/gin-gonic/gin"
)

func StartRouter(router *gin.Engine) {
	router.GET("/name", GetName)
	router.GET("/hobby", GetHobby)
	router.GET("/dream", GetDream)
}
