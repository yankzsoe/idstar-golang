package repositories

import (
	"errors"

	"gorm.io/gorm"
	config "idstar.com/session7/app/configs"
	"idstar.com/session7/app/models"
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

func (repo *UserRepository) FindByID(id string) (*models.UserModel, error) {
	var user models.UserModel
	result := repo.DB.Where("id = ?", id).Find(&user)
	if result.Error != nil {
		return nil, errors.New("failed to find user")
	}
	return &user, nil
}

func (repo *UserRepository) Update(user *models.UserModel) error {
	result := repo.DB.Save(user)
	if result.Error != nil {
		return errors.New("failed to update user")
	}
	return nil
}

func (repo *UserRepository) Delete(user *models.UserModel) error {
	result := repo.DB.Delete(user)
	if result.Error != nil {
		return errors.New("failed to delete user")
	}
	return nil
}
