package handler

import (
	"net/http"

	"github.com/IgorAfonso/GoOportunitiesAPI/schemas"
	"github.com/gin-gonic/gin"
)

func ListOpeningHendler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil{
		sendErr(ctx, http.StatusInternalServerError, "error listing openings")
		return
	}
	sendSucces(ctx, "lise-openings", openings)
}