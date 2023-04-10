package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"idstar.com/session7/app/dtos"
	"idstar.com/session7/app/models"
	"idstar.com/session7/app/services"
)

type UserController struct {
	service *services.UserService
}

func NewUserController(service *services.UserService) *UserController {
	return &UserController{
		service: service,
	}
}

// GetUser godoc
//	@Summary		Get User Data
//	@Description	Get Dummy User Account
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			id	path	dtos.GetUserByID	false	"ID"
//	@Router			/user/{id} [get]
func (ctrl *UserController) GetUser(ctx *gin.Context) {
	req := dtos.GetUserByID{
		Id: ctx.Param("id"),
	}
	err := req.Validate()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	result, err := ctrl.service.GetUserByID(req.Id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	user := dtos.CreateUserResponse{
		Id:          result.Id,
		Username:    result.Username,
		Nickname:    result.Nickname,
		Email:       result.Email,
		CreatedDate: result.CreatedDate,
		UpdatedDate: result.UpdatedDate,
	}

	ctx.JSON(http.StatusOK, gin.H{"data": user})
}

// PostUser godoc
//	@Summary		Post User Data
//	@Description	Add new fake User Account
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			request	body	dtos.CreateUserRequest	true	"User"
//	@Router			/user [post]
func (ctrl *UserController) PostUser(ctx *gin.Context) {
	req := dtos.CreateUserRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := req.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	m := models.UserModel{
		Username: req.Username,
		Nickname: req.Nickname,
		Email:    req.Email,
		Password: req.Password,
	}

	result, err := ctrl.service.CreateUser(&m)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	createdUser := dtos.CreateUserResponse{
		Id:          result.Id,
		Username:    result.Username,
		Nickname:    result.Nickname,
		Email:       result.Email,
		CreatedDate: result.CreatedDate,
		UpdatedDate: result.UpdatedDate,
	}

	ctx.JSON(http.StatusCreated, gin.H{"data": createdUser})
}
