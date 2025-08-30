package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOpeningHendler(ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H{
				"message": "POST Opening",
			})
}

func ShowOpeningHendler(ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H{
				"message": "GET Opening",
			})
}

func DeleteOpeningHendler(ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H{
				"message": "GET Opening",
			})
}

func UpdateOpeningHendler(ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H{
				"message": "GET Opening",
			})
}

func ListOpeningHendler(ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H{
				"message": "GET Opening",
			})
}