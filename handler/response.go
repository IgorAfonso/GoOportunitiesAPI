package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func sendErr(ctx *gin.Context, code int, msg string){
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(code, gin.H{
		"msg": msg,
		"errorCode": code,
	})
}

func sendSucces(ctx *gin.Context, op string, data interface{}){
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("operation %s successful", op),
		"data": data,
	})
}