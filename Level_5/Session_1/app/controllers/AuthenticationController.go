package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"idstar.com/session8/app/dtos"
	"idstar.com/session8/app/services"
)

type AuthenticationController struct {
	service *services.AuthenticationService
}

func NewAuthenticationController(service *services.AuthenticationService) *AuthenticationController {
	return &AuthenticationController{
		service: service,
	}
}

// Login godoc
//
//	@Summary		Login user
//	@Description	Login user to authorization
//	@Tags			Authentication
//	@Accept			json
//	@Produce		json
//	@Param			request	body	dtos.LoginRequest	true	"body"
//	@Router			/auth/login [post]
func (ctrl *AuthenticationController) Login(ctx *gin.Context) {
	req := dtos.LoginRequest{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	statusCode, err := ctrl.service.Login(req)
	if err != nil {
		ctx.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(statusCode, gin.H{"data": "Successfully"})
}
