package handler

import (
	"net/http"

	"github.com/IgorAfonso/GoOportunitiesAPI/schemas"
	"github.com/gin-gonic/gin"
)

func ShowOpeningHendler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == ""{
		sendErr(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	opening := schemas.Opening{}
	if err := db.First(&opening, id).Error;err != nil {
		sendErr(ctx, http.StatusNotFound, "opening not found")
		return
	}
	sendSucces(ctx, "shor-opening", opening)
}