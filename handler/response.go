package handler

import (
	"fmt"
	"net/http"

	"github.com/IgorAfonso/GoOportunitiesAPI/schemas"
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

type ErrorResponse struct {
	Message string 		`json:"message"`
	ErrorCode string 	`json:"errorCode"`
}

type CreateOpeningResponse struct{
	Message string 					`json:"message"`
	Data schemas.OpeningResponse 	`json:"data"`
}

type DeleteOpeningResponse struct{
	Message string 					`json:"message"`
	Data schemas.OpeningResponse 	`json:"data"`
}

type ShowOpeningResponse struct{
	Message string 					`json:"message"`
	Data schemas.OpeningResponse 	`json:"data"`
}

type ListOpeningsResponse struct{
	Message string 					`json:"message"`
	Data []schemas.OpeningResponse 	`json:"data"`
}

type UpdateOpeningResponse struct{
	Message string 					`json:"message"`
	Data schemas.OpeningResponse 	`json:"data"`
}