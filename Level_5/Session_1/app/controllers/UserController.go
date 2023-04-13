package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"idstar.com/session8/app/dtos"
	"idstar.com/session8/app/models"
	"idstar.com/session8/app/services"
	"idstar.com/session8/app/tools"
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
//
//	@Summary		Get User Data
//	@Description	Get User Account
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			id	path	string	false	"ID"
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

// GetAllUser godoc
//
//	@Summary		Get All User Data
//	@Description	Get All User Account
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			request	query	dtos.CommonParam	false	"param"
//	@Router			/user/ [get]
func (ctrl *UserController) GetAllUser(ctx *gin.Context) {
	conv := tools.Conversion{}

	req := dtos.CommonParam{
		Where:  ctx.Query("username"),
		Limit:  conv.StrToInt(ctx.Query("limit")),
		Offset: conv.StrToInt(ctx.Query("offset")),
	}

	result, err := ctrl.service.GetAllUser(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	users := []dtos.CreateUserResponse{}
	for _, user := range *result {
		users = append(users, dtos.CreateUserResponse{
			Id:          user.Id,
			Username:    user.Username,
			Nickname:    user.Nickname,
			Email:       user.Email,
			CreatedDate: user.CreatedDate,
			UpdatedDate: user.UpdatedDate,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{"data": users})
}

// PostUser godoc
//
//	@Summary		Post User Data
//	@Description	Add new fake User Account
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			request	body	dtos.CreateOrUpdateUserRequest	true	"User"
//	@Router			/user/ [post]
func (ctrl *UserController) PostUser(ctx *gin.Context) {
	req := dtos.CreateOrUpdateUserRequest{}
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

// PutUser godoc
//
//	@Summary		Put User Data
//	@Description	Update User Account
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			id		path	string							true	"User ID"
//	@Param			request	body	dtos.CreateOrUpdateUserRequest	true	"User"
//	@Router			/user/{id} [put]
func (ctrl *UserController) PutUser(ctx *gin.Context) {
	req := dtos.CreateOrUpdateUserRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var userId string = ctx.Param("id")
	if userId == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid User ID"})
	}

	if err := req.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := ctrl.service.UpdateUser(userId, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": "Update Successfully"})
}

// DeleteUser godoc
//
//	@Summary		Delete User Data
//	@Description	Delete User Account
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			id	path	string	true	"User ID"
//	@Router			/user/{id} [delete]
func (ctrl *UserController) DeleteUser(ctx *gin.Context) {
	var userId string = ctx.Param("id")
	if userId == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid User ID"})
	}

	err := ctrl.service.DeleteUser(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": "Update Successfully"})
}
