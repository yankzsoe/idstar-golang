package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (ctrl *UserController) GetUser(ctx *gin.Context) {
	id := ctx.Param("id")
	ctx.JSON(http.StatusOK, gin.H{"user_id": id})
}

func (ctrl *UserController) PostUser(ctx *gin.Context) {
	// Create new user
	ctx.JSON(http.StatusCreated, "OK")
}
