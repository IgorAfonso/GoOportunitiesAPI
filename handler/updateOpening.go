package handler

import (
	"net/http"

	"github.com/IgorAfonso/GoOportunitiesAPI/schemas"
	"github.com/gin-gonic/gin"
)

func UpdateOpeningHendler(ctx *gin.Context) {
	request := UpdateOpeningRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate();err != nil{
		logger.Errorf("validation error: %v", err.Error())
		sendErr(ctx, http.StatusBadRequest, err.Error())
		return
	}

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

	if(request.Role) != "" {opening.Role = request.Role}
	if(request.Company) != "" {opening.Company = request.Company}
	if(request.Location) != "" {opening.Location = request.Location}
	if(request.Remote) != nil {opening.Remote = *request.Remote}
	if(request.Link) != "" {opening.Link = request.Link}
	if(request.Salary) > 0 {opening.Salary = request.Salary}

	if err := db.Save(&opening).Error; err != nil {
		logger.Errorf("error update opening: %v", err.Error())
		sendErr(ctx, http.StatusInternalServerError, "error updating opening")
		return
	}
	sendSucces(ctx, "update-opening", opening)
}