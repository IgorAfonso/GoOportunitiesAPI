package handler

import (
	"fmt"
	"net/http"

	"github.com/IgorAfonso/GoOportunitiesAPI/schemas"
	"github.com/gin-gonic/gin"
)

func DeleteOpeningHendler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == ""{
		sendErr(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	opening := schemas.Opening{}
	if err := db.First(&opening, id).Error; err != nil {
		sendErr(ctx, http.StatusNotFound, fmt.Sprintf("opening with id: %s not found", id))
		return
	}

	if err := db.Delete(&opening).Error; err != nil {
		sendErr(ctx, http.StatusInternalServerError, 
			fmt.Sprintf("error deleting opening with id: %s", id))
			return
	}
	sendSucces(ctx, "delete-opening", opening)
}