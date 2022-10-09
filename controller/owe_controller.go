package controller

import (
	"net/http"
	"you-owe-me/repository"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

// OweController : represent the owe's controller contract
type OweController interface {
	GetOwe(*gin.Context)
}

type oweController struct {
	oweRepo repository.OweRepository
}

// NewOweController -> returns new owe controller
func NewOweController(repo repository.OweRepository) OweController {
	return oweController{
		oweRepo: repo,
	}
}

func (h oweController) GetOwe(ctx *gin.Context) {
	uuidStr := ctx.Param("uuid")

	intID, err := uuid.FromString(uuidStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	owe, err := h.oweRepo.GetOwe(intID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, owe)

}
