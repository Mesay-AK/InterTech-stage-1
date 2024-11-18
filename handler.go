package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func GetName(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Mesay Abebe")
}

func GetHobby(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"hobby": "Reading"})
}

func GetDream(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Procrastination is an arrogance that God will give you another chance to do tommorow what he has already given you the chance today!! ---> AVOID THIS ARROGANCE.")
}
