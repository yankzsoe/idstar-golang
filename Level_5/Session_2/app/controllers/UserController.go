package controllers

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"idstar.com/session9/app/dtos"
	"idstar.com/session9/app/models"
	"idstar.com/session9/app/services"
	"idstar.com/session9/app/tools"
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
//	@Security		ApiKeyAuth
//	@Param			id	path	string	false	"ID"
//	@Router			/user/{id} [get]
func (ctrl *UserController) GetUser(ctx *gin.Context) {
	if err := CheckToken(ctx); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

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
//	@Security		ApiKeyAuth
//	@Param			request	query	dtos.CommonParam	false	"param"
//	@Router			/user/ [get]
func (ctrl *UserController) GetAllUser(ctx *gin.Context) {
	if err := CheckToken(ctx); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	conv := tools.Conversion{}
	req := dtos.CommonParam{
		Where:  ctx.Query("username"),
		Limit:  conv.StrToInt(ctx.Query("limit")),
		Offset: conv.StrToInt(ctx.Query("offset")),
	}

	result, err := ctrl.service.GetAllUser(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
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
//	@Security		ApiKeyAuth
//	@Param			request	body	dtos.CreateOrUpdateUserRequest	true	"User"
//	@Router			/user/ [post]
func (ctrl *UserController) PostUser(ctx *gin.Context) {
	if err := CheckToken(ctx); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

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
		RoleId:   req.RoleId,
	}

	result, err := ctrl.service.CreateUser(&m)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	createdUser := dtos.CreateUserResponse{
		Id:          result.Id,
		Username:    result.Username,
		Nickname:    result.Nickname,
		Email:       result.Email,
		CreatedDate: result.CreatedDate,
		UpdatedDate: result.UpdatedDate,
		RoleId:      result.RoleId,
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
//	@Security		ApiKeyAuth
//	@Param			id		path	string							true	"User ID"
//	@Param			request	body	dtos.CreateOrUpdateUserRequest	true	"User"
//	@Router			/user/{id} [put]
func (ctrl *UserController) PutUser(ctx *gin.Context) {
	if err := CheckToken(ctx); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	req := dtos.CreateOrUpdateUserRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var userId string = ctx.Param("id")
	if userId == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid User ID"})
		return
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
//	@Security		ApiKeyAuth
//	@Param			id	path	string	true	"User ID"
//	@Router			/user/{id} [delete]
func (ctrl *UserController) DeleteUser(ctx *gin.Context) {
	if err := CheckToken(ctx); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	var userId string = ctx.Param("id")
	if userId == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid User ID"})
		return
	}

	err := ctrl.service.DeleteUser(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": "Update Successfully"})
}

const key = "abcdefghij1234567890"

func CheckToken(c *gin.Context) error {

	// Get token from header
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return errors.New("Missing Header 'Authorization'")
	}

	jwtKey := []byte(key)
	// Verify token
	tokenString := authHeader[len("Bearer "):]
	token, err := jwt.ParseWithClaims(tokenString, &dtos.Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrInvalidKey
		}
		return jwtKey, nil
	})

	if err != nil {
		return err
	}

	claims, ok := token.Claims.(*dtos.Claims)
	if !ok || !token.Valid {
		return errors.New("Not Authorize")
	}

	method := c.Request.Method
	url := c.Request.URL
	success := ClaimChecker(method, url.Path, *claims)

	if !success {
		return errors.New("Can't access this resources")
	}

	return nil
}

func ClaimChecker(method string, url string, claim dtos.Claims) bool {
	permissions := claim.Role.Permissions

	module := strings.Split(url, "/")[3]
	for _, permission := range permissions {
		if strings.EqualFold(permission.Module, module) {
			switch strings.ToLower(method) {
			case "post":
				return permission.CanCreate
			case "get":
				return permission.CanRead
			case "put":
				return permission.CanUpdate
			case "delete":
				return permission.CanDelete
			default:
				return false
			}
		}
	}

	return false
}
