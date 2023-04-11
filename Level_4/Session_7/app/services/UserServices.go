package services

import (
	"idstar.com/session7/app/dtos"
	"idstar.com/session7/app/models"
	"idstar.com/session7/app/repositories"
	"idstar.com/session7/app/tools"
)

type UserService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) *UserService {
	return &UserService{userRepository}
}

func (u *UserService) CreateUser(user *models.UserModel) (*models.UserModel, error) {
	// Enkripsi password
	aes128 := tools.Aes128{}
	encryptedPassword, err := aes128.Encrypt(user.Password)
	if err != nil {
		return nil, err
	}

	// Simpan user ke database
	user.Password = *encryptedPassword
	return u.userRepository.Create(user)
}

func (u *UserService) GetUserByID(userID string) (*models.UserModel, error) {
	return u.userRepository.FindByID(userID)
}

func (u *UserService) GetAllUser(param dtos.CommonParam) (*[]models.UserModel, error) {
	return u.userRepository.FindAll(param)
}

func (u *UserService) UpdateUser(userId string, user dtos.CreateOrUpdateUserRequest) error {
	return u.userRepository.Update(userId, user)
}

func (u *UserService) DeleteUser(userId string) error {
	return u.userRepository.Delete(userId)
}
