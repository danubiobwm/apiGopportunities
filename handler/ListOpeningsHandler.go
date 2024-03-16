package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/danubiobwm/apiGopportunities/schemas"
)

func ListOpeningsHandler(ctx *gin.Context) {

	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "Error fetching openings from database")
		return
	}

	sendSuccess(ctx, "list-openings", openings)
}
