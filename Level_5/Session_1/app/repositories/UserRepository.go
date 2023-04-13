package repositories

import (
	"errors"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	config "idstar.com/session8/app/configs"
	"idstar.com/session8/app/dtos"
	"idstar.com/session8/app/models"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		DB: config.GetDB(),
	}
}

func (repo *UserRepository) Create(user *models.UserModel) (*models.UserModel, error) {
	result := repo.DB.Create(user)
	if result.Error != nil {
		return nil, errors.New("failed to create user")
	}
	return user, nil
}

func (repo *UserRepository) FindAll(param dtos.CommonParam) (*[]models.UserModel, error) {
	var user []models.UserModel
	result := repo.DB.Where("username LIKE ?", "%"+param.Where+"%").Limit(param.Limit).Offset(param.Offset).Find(&user)
	if result.Error != nil {
		return nil, errors.New("failed to find user")
	}
	return &user, nil
}

func (repo *UserRepository) FindByID(id string) (*models.UserModel, error) {
	var user models.UserModel
	result := repo.DB.Where("id = ?", id).Find(&user)
	if result.Error != nil {
		return nil, errors.New("failed to find user")
	}
	return &user, nil
}

func (repo *UserRepository) Update(userId string, user dtos.CreateOrUpdateUserRequest) error {
	tNow := time.Now()
	if err := repo.DB.Model(models.UserModel{}).Where("id = ?", userId).Updates(models.UserModel{
		Username:    user.Username,
		Nickname:    user.Nickname,
		Email:       user.Email,
		Password:    user.ConfirmPassword,
		UpdatedDate: &tNow,
	}).Error; err != nil {
		return errors.New("failed to update user")
	}
	return nil
}

func (repo *UserRepository) Delete(userId string) error {
	user := models.UserModel{}
	if err := repo.DB.Clauses(clause.Returning{}).Delete(&user, "id", userId).Error; err != nil {
		return errors.New("failed to delete user")
	}

	if user.Id == "" {
		return errors.New("user not found")
	}

	return nil
}

func (repo *UserRepository) FindByUsernameOrEmail(username string) (*models.UserModel, error) {
	var user models.UserModel
	result := repo.DB.Where("username = ? OR email = ?", username, username).Find(&user)
	if result.Error != nil {
		return nil, errors.New("failed to find user")
	}
	return &user, nil
}
